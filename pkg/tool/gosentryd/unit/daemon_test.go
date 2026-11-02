package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosentryd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New()
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

func TestCloseClosesTheSentryClient(t *testing.T) {
	c := service_tester.NewSource(t)
	daemon_tester.New(t, c).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestWaitAnswersOnceSentryAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		daemon_tester.New(t, service_tester.NewSource(t)).Wait(t.Context()),
	)
}
