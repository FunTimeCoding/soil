package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	goraidparsedConstant "github.com/funtimecoding/soil/pkg/tool/goraidparsed/constant"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, service_tester.WithReporter(r))
	s.Parser.Fail(outage_tester.RefusedDial())

	return s.Served(), outage_tester.Guard(
		t,
		s.Daemon().Workers(),
		goraidparsedConstant.Identity.Name(),
	).Unavailable
}
