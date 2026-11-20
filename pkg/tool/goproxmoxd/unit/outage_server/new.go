package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(
		t,
		service_tester.WithNetbox(),
		service_tester.WithReporter(r),
	)

	switch peer {
	case gonetboxdConstant.Identity.Name():
		s.Netbox.Failure = outage_tester.RefusedDial()
	default:
		s.Alfa().Hypervisor.SetFailure(outage_tester.RefusedDial())
	}

	return s.Served(), outage_tester.New(t).Guard(s.Daemon().Workers(), peer).Unavailable
}
