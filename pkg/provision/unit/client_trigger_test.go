package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/provision/client"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestClientTriggerQueuesRunWithChanges(t *testing.T) {
	n := mock_runner.New()
	s := newTriggerServer(t, n, memory.New())
	c := client.New(
		webConstant.Localhost,
		s.Port,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	assert.FatalOnError(t, c.Trigger([]string{"alfa", "bravo"}))
	requests := n.Requests()
	assert.Count(t, 1, requests)
	assert.True(t, requests[0].Update)
	assert.Any(
		t,
		[]string{"alfa", "bravo"},
		requests[0].Parameters[constant.DownstreamChanges],
	)
}

func TestClientTriggerWithoutTokenIsRefused(t *testing.T) {
	n := mock_runner.New()
	s := newTriggerServer(t, n, memory.New())
	c := client.New(webConstant.Localhost, s.Port, "bravo")
	defer c.Close()
	e := c.Trigger([]string{"alfa"})
	assert.NotNil(t, e)
	assert.StringContains(t, "status 401", e.Error())
	assert.Count(t, 0, n.Requests())
}

func TestClientTriggerReportsQueuedRunAsConflict(t *testing.T) {
	r := memory.New()
	s := newTriggerServer(
		t,
		mock_runner.New().WithFail(conflict.Format("run already queued")),
		r,
	)
	c := client.New(
		webConstant.Localhost,
		s.Port,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	e := c.Trigger([]string{"alfa"})
	assert.NotNil(t, e)
	assert.StringContains(t, "status 409", e.Error())
	assert.StringContains(t, "run already queued", e.Error())
	assert.Count(t, 0, r.Events())
}

func TestClientTriggerFailureIsCaptured(t *testing.T) {
	r := memory.New()
	s := newTriggerServer(t, mock_runner.New().WithFail(fmt.Errorf("alfa")), r)
	c := client.New(
		webConstant.Localhost,
		s.Port,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	e := c.Trigger([]string{"alfa"})
	assert.NotNil(t, e)
	assert.StringContains(t, "status 500", e.Error())
	assert.Count(t, 1, r.Events())
}
