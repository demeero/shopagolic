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
	DBName        string `default:"shopagolic-catalog" split_words:"true"`
	MongoURI      string `default:"mongodb://localhost:27017" split_words:"true"`
	GRPCAddr      string `default:"localhost:8080" split_words:"true"`
	MigrationPath string `default:"../migrations" split_words:"true"`
}

func getConfig(t *testing.T) testConfig {
	t.Helper()
	once.Do(func() {
		err := envconfig.Process("", &cfg)
		require.NoError(t, err, "failed load integration test config")
	})
	return cfg
}
