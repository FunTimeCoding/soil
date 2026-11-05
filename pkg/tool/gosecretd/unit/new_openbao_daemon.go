package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/openbao/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/option"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func newOpenbaoDaemon(t *testing.T) *daemon.Daemon {
	t.Helper()
	c := mock_client.New()
	o := option.New()
	o.Source = func() face.OpenbaoClient { return c }
	result := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	t.Cleanup(result.Close)

	return result
}
