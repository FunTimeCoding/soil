package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/habitica/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/option"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New()
	o := option.New()
	o.Host = "habitica.test"
	o.Source = func() face.HabiticaClient { return c }
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

func TestCloseClosesHabiticaClient(t *testing.T) {
	c := service_tester.NewSource(t)
	daemon_tester.New(t, c).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestWaitAnswersOnceHabiticaAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		daemon_tester.New(t, service_tester.NewSource(t)).Wait(t.Context()),
	)
}
