package rpc

import (
	"context"
	"errors"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/demeero/shopagolic/cartservice/cart"
	"github.com/demeero/shopagolic/services/go/bricks/zaplogger"
	cartpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/cart/v1beta1"
)

type CartComponents struct {
	Adder   *cart.Adder
	Loader  *cart.Loader
	Deleter *cart.Deleter
}

type Cart struct {
	cartpb.UnimplementedCartServiceServer
	adder   *cart.Adder
	loader  *cart.Loader
	deleter *cart.Deleter
}

func NewCart(components CartComponents) *Cart {
	return &Cart{adder: components.Adder, loader: components.Loader, deleter: components.Deleter}
}

func (c *Cart) AddItem(ctx context.Context, req *cartpb.AddItemRequest) (*cartpb.AddItemResponse, error) {
	if err := errHandler(ctx, c.adder.AddItem(ctx, req.GetUserId(), convertProtoItem(req.GetItem()))); err != nil {
		return nil, err
	}
	return &cartpb.AddItemResponse{}, nil
}
func (c *Cart) GetCart(ctx context.Context, req *cartpb.GetCartRequest) (*cartpb.GetCartResponse, error) {
	loaded, err := c.loader.LoadByUserID(ctx, req.GetUserId())
	if err := errHandler(ctx, err); err != nil {
		return nil, err
	}
	return &cartpb.GetCartResponse{
		UserId: loaded.UserID,
		Items:  convertItems(loaded.Items),
	}, nil
}
func (c *Cart) EmptyCart(ctx context.Context, req *cartpb.EmptyCartRequest) (*cartpb.EmptyCartResponse, error) {
	if err := errHandler(ctx, c.deleter.DeleteAll(ctx, req.GetUserId())); err != nil {
		return nil, err
	}
	return &cartpb.EmptyCartResponse{}, nil
}

func errHandler(ctx context.Context, err error) error {
	if err == nil {
		return nil
	}
	if errors.Is(err, cart.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if errors.Is(err, cart.ErrInvalidData) {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	if errors.Is(err, cart.ErrConflictData) {
		return status.Error(codes.AlreadyExists, err.Error())
	}
	zaplogger.FromCtx(ctx).Error("failed handle request", zap.Error(err))
	return err
}

func convertProtoItem(item *cartpb.CartItem) cart.Item {
	return cart.Item{
		ProductID: item.GetProductId(),
		Quantity:  uint16(item.GetQuantity()),
	}
}

func convertItems(items []cart.Item) []*cartpb.CartItem {
	result := make([]*cartpb.CartItem, 0, len(items))
	for _, item := range items {
		result = append(result, convertItem(item))
	}
	return result
}

func convertItem(item cart.Item) *cartpb.CartItem {
	return &cartpb.CartItem{
		ProductId: item.ProductID,
		Quantity:  int32(item.Quantity),
	}
}
