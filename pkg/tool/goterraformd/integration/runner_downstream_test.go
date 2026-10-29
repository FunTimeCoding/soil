package integration

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	"github.com/funtimecoding/soil/pkg/provision"
	provisionConstant "github.com/funtimecoding/soil/pkg/provision/constant"
	"github.com/funtimecoding/soil/pkg/provision/mock_runner"
	"github.com/funtimecoding/soil/pkg/provision/types/host_event"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/terraform_tester"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net"
	"net/http"
	"strconv"
	"testing"
)

func TestApplySendsHostEventsDownstream(t *testing.T) {
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
	terraform_tester.NewWithDownstream(
		t,
		mock_client.New(),
		constant.FixturePlan,
		net.JoinHostPort(webConstant.Localhost, strconv.Itoa(v.Port)),
	)
	requests := waitForRequest(t, n)
	assert.Any(
		t,
		[]*host_event.Event{
			host_event.New("alfa", "update"),
			host_event.New("bravo", "update"),
			host_event.New("foxtrot", "delete"),
			host_event.New("golf", "replace"),
			host_event.New("hotel", "update"),
			host_event.New("lima", "update"),
		},
		requests[0].Parameters[provisionConstant.DownstreamHosts],
	)
	assert.Count(
		t,
		8,
		requests[0].Parameters[provisionConstant.DownstreamChanges],
	)
}
