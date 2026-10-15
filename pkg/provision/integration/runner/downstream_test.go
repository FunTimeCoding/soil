package runner

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/provision"
	"github.com/funtimecoding/soil/pkg/provision/client"
	"github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	"github.com/funtimecoding/soil/pkg/provision/types/trigger"
	"github.com/funtimecoding/soil/pkg/provision/unit/runner_tester"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestApplyWithChangesTriggersDownstream(t *testing.T) {
	n := mock_runner.New()
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			provision.Mount(n, memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	c := client.New(
		webConstant.Localhost,
		v.Port,
		generativeConstant.ModelContextTestToken,
	)
	defer c.Close()
	s := runner_tester.NewDownstream(t, c)
	response := make(chan *trigger.Result, 1)
	s.Trigger(trigger.Request{Response: response})
	<-response
	requests := n.Requests()
	assert.Count(t, 2, requests)
	assert.Any(
		t,
		[]string{"alfa"},
		requests[1].Parameters[constant.DownstreamChanges],
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestRefusedDownstreamIsCaptured(t *testing.T) {
	n := mock_runner.New()
	v := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			provision.Mount(n, memory.New(), mock_recorder.New(), g)
		},
	)
	defer v.Close()
	c := client.New(webConstant.Localhost, v.Port, "bravo")
	defer c.Close()
	s := runner_tester.NewDownstream(t, c)
	response := make(chan *trigger.Result, 1)
	s.Trigger(trigger.Request{Response: response})
	<-response
	assert.Count(t, 0, n.Requests())
	events := s.Reporter.Events()
	assert.Count(t, 2, events)
	assert.StringContains(t, "status 401", events[0].Error.Error())
	assert.Any(
		t,
		[]string{"alfa"},
		events[0].Context[constant.DownstreamChanges],
	)
}
