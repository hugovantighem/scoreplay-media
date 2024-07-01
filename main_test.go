package main_test

import (
	"context"
	"myproject/common"
	"myproject/infra"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// Tests the application starts
func TestApplication(t *testing.T) {
	conf := infra.Config{
		ServerAddr:   "0.0.0.0:8080",
		DbConnString: "mongodb://localhost:27017/default",
	}
	stop := infra.RunApplication(conf)
	defer stop()

	result, err := common.WithRetry(context.Background(),
		func(ctx context.Context) (any, error) {
			result, err := http.Get("http://" + conf.ServerAddr + "/ping")
			if err != nil {
				return nil, err
			}

			return result, nil
		}, func(_ context.Context, _ interface{}, err error) (bool, error) {
			if err != nil {
				return true, err
			}

			return false, nil
		}, 10, 100*time.Millisecond)
	require.NoError(t, err)
	res, ok := result.(*http.Response)
	require.True(t, ok)
	require.Equal(t, http.StatusOK, res.StatusCode)
}
