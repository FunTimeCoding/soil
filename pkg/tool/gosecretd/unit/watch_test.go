package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/watch"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func TestWatchReportsAnUnsealedStore(t *testing.T) {
	y := prometheus.NewRegistry()
	watch.New(
		mock_client.New(),
		y,
		logger.New(context.Background()),
		memory.New(),
	).Poll()
	assert.Float(t, 0, gauge(t, y, constant.SealedGauge))
	assert.Float(t, 1, gauge(t, y, constant.ReachableGauge))
}

func TestWatchReportsASealedStore(t *testing.T) {
	y := prometheus.NewRegistry()
	c := mock_client.New()
	c.Seal()
	watch.New(c, y, logger.New(context.Background()), memory.New()).Poll()
	assert.Float(t, 1, gauge(t, y, constant.SealedGauge))
}

func TestWatchReportsAnUnreachableStore(t *testing.T) {
	y := prometheus.NewRegistry()
	c := mock_client.New()
	c.FailHealth(errors.New("connection refused"))
	watch.New(c, y, logger.New(context.Background()), memory.New()).Poll()
	assert.Float(t, 0, gauge(t, y, constant.ReachableGauge))
}
