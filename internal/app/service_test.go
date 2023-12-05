package app_test

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/tellmeac/goalgo/internal/app"
	"testing"
	"time"
)

func newApp() *app.Service {
	return app.New(&app.Config{
		DatabaseConn: "./test.db",
	})
}

func TestService_GetLatest(t *testing.T) {
	a := newApp()

	chart, err := a.GetLatest(
		context.Background(),
		time.Date(2022, 5, 25, 0, 0, 0, 0, time.UTC).Unix(),
	)

	require.NoError(t, err)
	require.NotEmpty(t, chart.Data)
}

func TestService_GetInPeriod(t *testing.T) {
	a := newApp()

	chart, err := a.GetInPeriod(
		context.Background(),
		time.Date(2022, 5, 25, 0, 0, 0, 0, time.UTC).Unix(),
		time.Date(2023, 5, 30, 0, 0, 0, 0, time.UTC).Unix(),
	)

	require.NoError(t, err)
	require.NotEmpty(t, chart.Data)
}
