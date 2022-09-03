package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	cartpb "github.com/demeero/shopagolic/services/proto/gen/go/shopagolic/cart/v1beta1"
)

func TestIntegrationHealth(t *testing.T) {
	t.Logf("health test config: %+v", getConfig(t))
	_, err := healthGRPCClient(t).Health(context.Background(), &cartpb.HealthRequest{})
	assert.NoError(t, err)
}

func healthGRPCClient(t *testing.T) cartpb.HealthServiceClient {
	t.Helper()
	conn, err := grpc.Dial(getConfig(t).GRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	return cartpb.NewHealthServiceClient(conn)
}
