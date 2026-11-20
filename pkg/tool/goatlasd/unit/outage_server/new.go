package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/base"
	gonetboxdConstant "github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	gooutpostdConstant "github.com/funtimecoding/soil/pkg/tool/gooutpostd/constant"
	gopnsensedConstant "github.com/funtimecoding/soil/pkg/tool/gopnsensed/constant"
	goprocessdConstant "github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, base.WithReporter(r), base.WithOutpost())

	switch peer {
	case gonetboxdConstant.Identity.Name():
		s.Netbox.Failure = outage_tester.RefusedDial()
	case constant.SourceKubernetes:
		s.Kubernetes.Failure = outage_tester.RefusedDial()
	case goprocessdConstant.Identity.Name():
		s.Process.Failure = outage_tester.RefusedDial()
	case join.Colon(
		gooutpostdConstant.Identity.Name(),
		constant.FixtureVirtualMachineNode,
	):
		s.Outpost.Failure = outage_tester.RefusedDial()
	case gopnsensedConstant.Identity.Name():
		s.Lease.Failure = outage_tester.RefusedDial()
	}

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		peer,
	).Unavailable
}
