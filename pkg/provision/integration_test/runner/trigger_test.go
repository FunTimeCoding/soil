package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/integration_test/runner_tester"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"testing"
)

func TestTriggerCallsApply(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.Trigger(runner.TriggerRequest{})
	s.WaitForApply(2)
	assert.String(t, "manual", s.LastApply().TriggerSource)
}

func TestTriggerWithParameters(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	s.Trigger(
		runner.TriggerRequest{
			Parameters: map[string]any{"target": "specific"},
		},
	)
	s.WaitForApply(2)
	assert.String(t, "specific", s.LastApply().Parameters["target"].(string))
}

func TestSynchronousTrigger(t *testing.T) {
	s := runner_tester.New(t)
	s.WaitForApply(1)
	response := make(chan *runner.TriggerResult, 1)
	s.Trigger(
		runner.TriggerRequest{
			Response: response,
		},
	)
	result := <-response
	assert.Nil(t, result.Error)
	assert.NotNil(t, result.Value)
}
