package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	currpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/currency/v1beta1"
)

func TestIntegrationHealth(t *testing.T) {
	t.Logf("currency test config: %+v", getConfig(t))
	_, err := healthGRPCClient(t).Health(context.Background(), &currpb.HealthRequest{})
	assert.NoError(t, err)
}

func healthGRPCClient(t *testing.T) currpb.HealthServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return currpb.NewHealthServiceClient(conn)
}
