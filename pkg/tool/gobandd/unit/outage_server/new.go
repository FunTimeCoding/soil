package outage_server

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
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
	s.Alfa().SetPower(
		bandConstant.PowerOn,
		bandConstant.PowerOn,
		bandConstant.PowerOffSoftGraceful,
	)

	switch peer {
	case gonetboxdConstant.Identity.Name():
		s.Netbox.Failure = outage_tester.RefusedDial()
	default:
		s.Alfa().Fail(outage_tester.RefusedDial())
	}

	return s.Served(), outage_tester.Guard(t, s.Daemon().Workers(), peer).Unavailable
}
