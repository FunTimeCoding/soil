package unit

import (
	"context"
	"errors"
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/provision/client"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	generated "github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/web"
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
	assert.FatalOnError(
		t,
		c.PostTrigger(
			change.New(
				[]string{"alfa", "bravo"},
				[]*host_event.Event{
					host_event.New("charlie", constant.HostReplace),
					host_event.New("delta", constant.HostDelete),
				},
			),
		),
	)
	requests := n.Requests()
	assert.Count(t, 1, requests)
	assert.True(t, requests[0].Update)
	assert.Any(
		t,
		[]string{"alfa", "bravo"},
		requests[0].Parameters[constant.DownstreamChanges],
	)
	assert.Any(
		t,
		[]*host_event.Event{
			host_event.New("charlie", "replace"),
			host_event.New("delta", "delete"),
		},
		requests[0].Parameters[constant.DownstreamHosts],
	)
}

func TestClientTriggerReachingNoHostSendsEmptyHosts(t *testing.T) {
	n := mock_runner.New()
	s := newTriggerServer(t, n, memory.New())
	c := client.New(
		webConstant.Localhost,
		s.Port,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	assert.FatalOnError(t, c.PostTrigger(change.New([]string{"alfa"}, nil)))
	assert.Any(
		t,
		[]*host_event.Event{},
		n.Requests()[0].Parameters[constant.DownstreamHosts],
	)
}

func TestTriggerWithoutHostsFieldCarriesNoHosts(t *testing.T) {
	n := mock_runner.New()
	s := newTriggerServer(t, n, memory.New())
	c, e := generated.NewClientWithResponses(
		s.Address(""),
		generated.WithHTTPClient(s.Client()),
		generated.WithRequestEditorFn(
			web.BearerEditor(generativeConstant.ModelContextTestToken),
		),
	)
	assert.FatalOnError(t, e)
	v, f := c.PostTriggerWithResponse(
		context.Background(),
		generated.TriggerRequest{Update: true, Changes: []string{"alfa"}},
	)
	assert.FatalOnError(t, f)
	assert.Integer(t, 202, v.StatusCode())
	_, present := n.Requests()[0].Parameters[constant.DownstreamHosts]
	assert.False(t, present)
}

func TestClientTriggerWithoutTokenIsRefused(t *testing.T) {
	n := mock_runner.New()
	s := newTriggerServer(t, n, memory.New())
	c := client.New(webConstant.Localhost, s.Port, "bravo")
	defer c.Close()
	e := c.PostTrigger(change.New([]string{"alfa"}, nil))
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "401", e.Error())
	assert.Count(t, 0, n.Requests())
}

func TestClientTriggerToNobodyIsAConnectionFailure(t *testing.T) {
	c := client.New(
		webConstant.Localhost,
		1,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	e := c.PostTrigger(change.New([]string{"alfa"}, nil))
	var f *connection.Failure
	assert.True(t, errors.As(e, &f))
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
	e := c.PostTrigger(change.New([]string{"alfa"}, nil))
	assert.True(t, conflict.Is(e))
	assert.StringContains(t, "run already queued", e.Error())
	assert.Count(t, 0, r.Events())
}

func TestClientTriggerOnAStoppedRunnerIsUnavailable(t *testing.T) {
	r := memory.New()
	s := newTriggerServer(
		t,
		mock_runner.New().WithFail(unavailable.Format(constant.RunnerStopped)),
		r,
	)
	c := client.New(
		webConstant.Localhost,
		s.Port,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	e := c.PostTrigger(change.New([]string{"alfa"}, nil))
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "runner stopped", e.Error())
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
	e := c.PostTrigger(change.New([]string{"alfa"}, nil))
	assert.True(t, unexpected.Is(e))
	assert.StringContains(t, "500", e.Error())
	assert.Count(t, 1, r.Events())
}
