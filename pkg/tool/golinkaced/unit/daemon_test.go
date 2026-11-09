package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/linkace/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/daemon"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/option"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestCloseClosesLinkAceClient(t *testing.T) {
	c := service_tester.NewSource(t)
	daemon_tester.New(t, c, memory.New()).Close()
	assert.Integer(t, 1, c.Closed)
}

func TestWaitAnswersOnceLinkAceAnswers(t *testing.T) {
	assert.FatalOnError(
		t,
		daemon_tester.New(t, service_tester.NewSource(t), memory.New()).
			Wait(t.Context()),
	)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	c := mock_client.New()
	o := option.New()
	o.Source = func() face.LinkAceClient { return c }
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
