package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/kubernetes/client/mock_client"
	provisionMock "github.com/funtimecoding/soil/pkg/provision/client/mock_client"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	downstream := provisionMock.New(t, "alfa.example")
	downstream.Fail(downstream.Reachable, outage_tester.RefusedDial())
	downstream.Fail(downstream.PostTrigger, outage_tester.RefusedDial())
	s := base.New(
		t,
		mock_client.New(),
		service_tester.WithReporter(r),
		service_tester.WithDownstream(downstream),
	)
	workers := s.Daemon().Workers()
	workers[len(workers)-1].Close()

	return s.Served(), outage_tester.New(t).Guard(
		workers,
		join.Colon(constant.DownstreamPeer, downstream.Address()),
	).Unavailable
}
