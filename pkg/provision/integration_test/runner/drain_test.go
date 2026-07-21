package runner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"testing"
	"time"
)

func TestDrainChannelsOnStop(t *testing.T) {
	gate := make(chan struct{})
	t.Cleanup(
		func() {
			select {
			case <-gate:
			default:
				close(gate)
			}
		},
	)
	r := runner.New(
		runner.Configuration{
			SetupFunction: func() bool {
				<-gate

				return false
			},
			ApplyFunction: func(
				_ map[string]any,
				_ string,
			) any {
				return nil
			},
		},
		logger.New(context.Background()),
		memory.New(),
	)
	r.Start()
	response := make(chan *runner.TriggerResult, 1)
	e := r.Trigger(runner.TriggerRequest{Response: response})
	assert.FatalOnError(t, e)
	close(gate)

	select {
	case result := <-response:
		assert.Error(t, result.Error)
	case <-time.After(5 * time.Second):
		t.Fatal("timed out waiting for drain response")
	}
}
