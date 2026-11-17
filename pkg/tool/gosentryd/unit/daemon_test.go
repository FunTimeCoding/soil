package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/base"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New(t)
	o := option.New()
	o.Organization = "test"
	o.Host = "sentry.test"
	o.Source = func() face.SentryClient { return c }
	d := daemon.New(o, logger.New(t.Context()), memory.New())
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

func TestCloseClosesSentryClient(t *testing.T) {
	s := base.New(t)
	s.Daemon().Close()
	assert.Integer(t, 1, s.Source.Closed)
}

func TestWaitAnswersOnceSentryAnswers(t *testing.T) {
	assert.FatalOnError(t, base.New(t).Daemon().Wait(t.Context()))
}
