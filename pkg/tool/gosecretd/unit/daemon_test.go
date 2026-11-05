package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/openbao/face"
	"github.com/funtimecoding/soil/pkg/openbao/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/option"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	c := mock_client.New()
	o := option.New()
	o.Source = func() face.OpenbaoClient { return c }
	before := goleak.IgnoreCurrent()
	d := daemon.New(
		o,
		logger.New(t.Context()),
		memory.New(),
		prometheus.NewRegistry(),
	)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}
