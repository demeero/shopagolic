package test

import (
	"sync"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"
)

var (
	cfg  testConfig
	once = &sync.Once{}
)

type testConfig struct {
	RedisAddr     string `default:"localhost:6379" split_words:"true"`
	CartKeyPrefix string `default:"cart:" split_words:"true"`
	GRPCAddr      string `default:"localhost:8080" split_words:"true"`
}

func getConfig(t *testing.T) testConfig {
	t.Helper()
	once.Do(func() {
		err := envconfig.Process("", &cfg)
		require.NoError(t, err, "failed load integration test config")
	})
	return cfg
}
