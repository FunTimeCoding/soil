package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func TestReaperInterruptStopsOwnedProcess(t *testing.T) {
	windowsSkip(t)
	owned := reporter.NewOptional(identity.Example().Name())
	defer owned.Close()
	r := reaper.New(owned, prometheus.NewRegistry())
	defer r.Close()
	c := run.New().NoPanic().WithRegistry(r)
	finished := make(chan struct{})
	go func() {
		c.Start("sleep", "30")
		close(finished)
	}()
	timeout := time.After(5 * time.Second)

	for {
		r.Interrupt()

		select {
		case <-finished:
			assert.StringContains(t, "signal: interrupt", c.Error.Error())

			return
		case <-timeout:
			t.Fatal("owned process outlived the interrupt")
		case <-time.After(50 * time.Millisecond):
		}
	}
}
