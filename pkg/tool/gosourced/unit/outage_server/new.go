package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	gomemorydConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/base"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, base.WithReporter(r))
	s.Memory.Failure = outage_tester.RefusedDial()

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		gomemorydConstant.Identity.Name(),
	).Unavailable
}
