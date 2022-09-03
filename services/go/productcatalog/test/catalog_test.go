package test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	moneypb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/money/v1"
	catalogpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/productcatalog/v1beta1"
)

var predefinedCategories = []string{
	"accessories",
	"clothing",
	"footwear",
	"hair",
	"beauty",
	"decor",
	"home",
	"kitchen",
}

type moneyDoc struct {
	CurrencyCode string `bson:"currency_code"`
	Units        int64  `bson:"units"`
	Nanos        int32  `bson:"nanos"`
}

type productDoc struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	Picture     string             `bson:"picture"`
	Price       moneyDoc
	Categories  []string  `bson:"categories"`
	CreatedAt   time.Time `bson:"created_at"`
}

func TestIntegrationCatalog(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	suite.Run(t, &catalogTestSuite{})
}

type catalogTestSuite struct {
	suite.Suite
	ctx        context.Context
	db         *mongo.Database
	coll       *mongo.Collection
	data       []productDoc
	grpcClient catalogpb.ProductCatalogServiceClient
}

func (ts *catalogTestSuite) SetupSuite() {
	ts.ctx = context.Background()
	ts.T().Logf("catalog test config: %+v", getConfig(ts.T()))
	ts.db = dbConnect(ts.T())
	ts.coll = ts.db.Collection("products")
	ts.grpcClient = catalogGRPCClient(ts.T())
}

func (ts *catalogTestSuite) TearDownSuite() {
	ts.Require().NoError(ts.db.Client().Disconnect(ts.ctx))
}

func (ts *catalogTestSuite) SetupTest() {
	n := 10
	docs := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		gofakeit.ShuffleStrings(predefinedCategories)
		doc := productDoc{
			ID:          primitive.NewObjectID(),
			Name:        gofakeit.LetterN(5),
			Description: gofakeit.HipsterSentence(3),
			Picture:     gofakeit.URL(),
			Price: moneyDoc{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			CreatedAt:  gofakeit.Date(),
		}
		docs = append(docs, doc)
		ts.data = append(ts.data, doc)
	}

	_, err := ts.coll.InsertMany(ts.ctx, docs)
	ts.Require().NoError(err)
}

func (ts *catalogTestSuite) TearDownTest() {
	ts.data = nil
	_, err := ts.coll.DeleteMany(ts.ctx, bson.M{})
	ts.Require().NoError(err)
}

func (ts *catalogTestSuite) TestCreate() {
	resp, err := ts.grpcClient.CreateProduct(context.Background(), &catalogpb.CreateProductRequest{
		Name:        gofakeit.LetterN(5),
		Description: gofakeit.HipsterSentence(3),
		Picture:     gofakeit.URL(),
		Price: &moneypb.Money{
			CurrencyCode: gofakeit.CurrencyShort(),
			Units:        int64(gofakeit.Number(0, 1000)),
			Nanos:        int32(gofakeit.Number(0, 1000)),
		},
		Categories: predefinedCategories[:gofakeit.Number(1, 4)],
	})
	ts.NoError(err)
	ts.NotEmpty(resp.GetId())

	docID, err := primitive.ObjectIDFromHex(resp.GetId())
	ts.Require().NoError(err)
	ts.NoError(ts.coll.FindOne(context.Background(), bson.M{"_id": docID}).Err())
}

func (ts *catalogTestSuite) TestCreate_InvalidInput() {
	tests := []struct {
		name string
		req  *catalogpb.CreateProductRequest
	}{
		{
			name: "name too long",
			req: &catalogpb.CreateProductRequest{
				Name:        gofakeit.LetterN(300),
				Description: gofakeit.HipsterSentence(3),
				Picture:     gofakeit.URL(),
				Price: &moneypb.Money{
					CurrencyCode: gofakeit.CurrencyShort(),
					Units:        int64(gofakeit.Number(0, 1000)),
					Nanos:        int32(gofakeit.Number(0, 1000)),
				},
				Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			},
		},
		{
			name: "name too short",
			req: &catalogpb.CreateProductRequest{
				Name:        gofakeit.Letter(),
				Description: gofakeit.HipsterSentence(3),
				Picture:     gofakeit.URL(),
				Price: &moneypb.Money{
					CurrencyCode: gofakeit.CurrencyShort(),
					Units:        int64(gofakeit.Number(0, 1000)),
					Nanos:        int32(gofakeit.Number(0, 1000)),
				},
				Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			},
		},
		{
			name: "name empty",
			req: &catalogpb.CreateProductRequest{
				Description: gofakeit.HipsterSentence(3),
				Picture:     gofakeit.URL(),
				Price: &moneypb.Money{
					CurrencyCode: gofakeit.CurrencyShort(),
					Units:        int64(gofakeit.Number(0, 1000)),
					Nanos:        int32(gofakeit.Number(0, 1000)),
				},
				Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			},
		},
		{
			name: "picture is not url",
			req: &catalogpb.CreateProductRequest{
				Name:        gofakeit.LetterN(11),
				Description: gofakeit.HipsterSentence(3),
				Picture:     gofakeit.Word(),
				Price: &moneypb.Money{
					CurrencyCode: gofakeit.CurrencyShort(),
					Units:        int64(gofakeit.Number(0, 1000)),
					Nanos:        int32(gofakeit.Number(0, 1000)),
				},
				Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			},
		},
		{
			name: "description too long",
			req: &catalogpb.CreateProductRequest{
				Name:        gofakeit.LetterN(11),
				Description: gofakeit.LetterN(5000),
				Picture:     gofakeit.URL(),
				Price: &moneypb.Money{
					CurrencyCode: gofakeit.CurrencyShort(),
					Units:        int64(gofakeit.Number(0, 1000)),
					Nanos:        int32(gofakeit.Number(0, 1000)),
				},
				Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			},
		},
		{
			name: "description too short",
			req: &catalogpb.CreateProductRequest{
				Name:        gofakeit.LetterN(11),
				Description: gofakeit.Letter(),
				Picture:     gofakeit.URL(),
				Price: &moneypb.Money{
					CurrencyCode: gofakeit.CurrencyShort(),
					Units:        int64(gofakeit.Number(0, 1000)),
					Nanos:        int32(gofakeit.Number(0, 1000)),
				},
				Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.CreateProduct(context.Background(), tt.req)
			ts.Nil(resp)
			ts.Error(err)
			s, _ := status.FromError(err)
			ts.Equal("InvalidArgument", s.Code().String())
		})
	}
}

func (ts *catalogTestSuite) TestLoadByID() {
	for _, item := range ts.data {
		expected := &catalogpb.GetProductResponse{Product: &catalogpb.Product{
			Id:          item.ID.Hex(),
			Name:        item.Name,
			Description: item.Description,
			Picture:     item.Picture,
			Price: &moneypb.Money{
				CurrencyCode: item.Price.CurrencyCode,
				Units:        item.Price.Units,
				Nanos:        item.Price.Nanos,
			},
			Categories: item.Categories,
			CreatedAt:  timestamppb.New(primitive.NewDateTimeFromTime(item.CreatedAt).Time()),
		}}

		actual, err := ts.grpcClient.GetProduct(context.Background(), &catalogpb.GetProductRequest{Id: item.ID.Hex()})
		ts.Equal(expected.String(), actual.String())
		ts.NoError(err)
	}
}

func (ts *catalogTestSuite) TestLoadByID_Error() {
	tests := []struct {
		name         string
		id           string
		expectedCode codes.Code
	}{
		{
			name:         "not found",
			id:           primitive.NewObjectID().Hex(),
			expectedCode: codes.NotFound,
		},
		{
			name:         "invalid data",
			id:           "incorrect_id_format",
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "empty id",
			expectedCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.grpcClient.GetProduct(ts.ctx, &catalogpb.GetProductRequest{Id: tt.id})
			s, _ := status.FromError(err)
			ts.Nil(actual)
			ts.Error(err)
			ts.Equal(tt.expectedCode.String(), s.Code().String())
		})
	}
}

func (ts *catalogTestSuite) TestLoad_Pagination() {
	var (
		pSize       = 3
		total       = len(ts.data)
		pToken      string
		expectedIDs = make([]string, 0, len(ts.data))
		actualIDs   = make([]string, 0, len(ts.data))
	)

	for _, item := range ts.data {
		expectedIDs = append(expectedIDs, item.ID.Hex())
	}

	for total > 0 {
		resp, err := ts.grpcClient.ListProducts(ts.ctx, &catalogpb.ListProductsRequest{PageSize: int32(pSize), PageToken: pToken})
		ts.Require().NoError(err)
		ts.Require().NotEmpty(resp.GetProducts())

		for _, p := range resp.GetProducts() {
			actualIDs = append(actualIDs, p.GetId())
		}

		total -= len(resp.GetProducts())

		pToken = resp.GetNextPageToken()
		if pToken != "" {
			ts.Equal(pSize, len(resp.GetProducts()))
		} else {
			ts.LessOrEqual(len(resp.GetProducts()), pSize)
		}
	}

	ts.Empty(pToken)
	ts.Equal(total, 0)
	ts.ElementsMatch(expectedIDs, actualIDs)
}

func (ts *catalogTestSuite) TestSearch() {
	var (
		doc1 = productDoc{
			ID:          primitive.NewObjectID(),
			Name:        "Java Hut",
			Description: "Coffee and cakes",
			Picture:     gofakeit.URL(),
			Price: moneyDoc{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			CreatedAt:  gofakeit.Date(),
		}
		doc2 = productDoc{
			ID:          primitive.NewObjectID(),
			Name:        "Burger Buns",
			Description: "Gourmet hamburgers Discount",
			Picture:     gofakeit.URL(),
			Price: moneyDoc{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			CreatedAt:  gofakeit.Date(),
		}
		doc3 = productDoc{
			ID:          primitive.NewObjectID(),
			Name:        "Coffee Shop",
			Description: "Just coffee",
			Picture:     gofakeit.URL(),
			Price: moneyDoc{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			CreatedAt:  gofakeit.Date(),
		}
		doc4 = productDoc{
			ID:          primitive.NewObjectID(),
			Name:        "Clothes Clothes Clothes",
			Description: "Discount clothing",
			Picture:     gofakeit.URL(),
			Price: moneyDoc{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			CreatedAt:  gofakeit.Date(),
		}
		doc5 = productDoc{
			ID:          primitive.NewObjectID(),
			Name:        "Java Shopping",
			Description: "Indonesian goods",
			Picture:     gofakeit.URL(),
			Price: moneyDoc{
				CurrencyCode: gofakeit.CurrencyShort(),
				Units:        int64(gofakeit.Number(0, 1000)),
				Nanos:        int32(gofakeit.Number(0, 1000)),
			},
			Categories: predefinedCategories[:gofakeit.Number(1, 4)],
			CreatedAt:  gofakeit.Date(),
		}
	)

	_, err := ts.coll.InsertMany(ts.ctx, []interface{}{doc1, doc2, doc3, doc4, doc5})
	ts.Require().NoError(err)

	tests := []struct {
		name           string
		keywords       string
		expectedDocIDs []string
	}{
		{
			name:           "coffee and shop - 3 docs",
			keywords:       "coffee shop",
			expectedDocIDs: []string{doc1.ID.Hex(), doc3.ID.Hex(), doc5.ID.Hex()},
		},
		{
			name:           "discount - 2 docs",
			keywords:       "discount",
			expectedDocIDs: []string{doc2.ID.Hex(), doc4.ID.Hex()},
		},
		{
			name:           "exact search - 1 docs",
			keywords:       "Burger Buns",
			expectedDocIDs: []string{doc2.ID.Hex()},
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			resp, err := ts.grpcClient.SearchProducts(ts.ctx, &catalogpb.SearchProductsRequest{Query: tt.keywords})
			ts.NoError(err)
			foundIDs := make([]string, 0, len(tt.expectedDocIDs))
			for _, p := range resp.GetProducts() {
				foundIDs = append(foundIDs, p.GetId())
			}
			ts.ElementsMatch(tt.expectedDocIDs, foundIDs)
		})
	}
}

func (ts *catalogTestSuite) TestSearch_Error() {
	tests := []struct {
		name         string
		query        string
		expectedCode codes.Code
	}{
		{
			name:         "empty query",
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "query too long",
			query:        gofakeit.LetterN(130),
			expectedCode: codes.InvalidArgument,
		},
		{
			name:         "query too short",
			query:        gofakeit.Letter(),
			expectedCode: codes.InvalidArgument,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			actual, err := ts.grpcClient.SearchProducts(ts.ctx, &catalogpb.SearchProductsRequest{Query: tt.query})
			s, _ := status.FromError(err)
			ts.Nil(actual)
			ts.Error(err)
			ts.Equal(tt.expectedCode.String(), s.Code().String())
		})
	}
}

func dbConnect(t *testing.T) *mongo.Database {
	t.Helper()
	c := getConfig(t)
	opts := &options.ClientOptions{}
	client, err := mongo.NewClient(opts.ApplyURI(c.MongoURI))
	require.NoError(t, err)
	require.NoError(t, client.Connect(context.Background()))
	require.NoError(t, client.Ping(context.Background(), nil))
	driver, err := mongodb.WithInstance(client, &mongodb.Config{DatabaseName: c.DBName})
	require.NoError(t, err)
	m, err := migrate.NewWithDatabaseInstance("file://"+c.MigrationPath, c.DBName, driver)
	require.NoError(t, err)
	err = m.Up()
	if !errors.Is(err, migrate.ErrNoChange) {
		require.NoError(t, err)
	}
	return client.Database(c.DBName)
}

func catalogGRPCClient(t *testing.T) catalogpb.ProductCatalogServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return catalogpb.NewProductCatalogServiceClient(conn)
}
