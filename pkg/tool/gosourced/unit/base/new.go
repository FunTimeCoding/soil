package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/source/inventory/module"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/unit/daemon_tester"
	"testing"
)

func New(
	t *testing.T,
	fixture string,
) *Server {
	t.Helper()
	d := testutil.PrepareTestPackage(t, fixture)
	i := inventory.New()
	i.Modules = []module.Module{{Name: "test", Directory: d}}
	owner, _ := daemon_tester.New(t, i)
	result := &Server{
		Server:    model_context_server.Mount(t, owner, memory.New()),
		Directory: d,
	}
	t.Cleanup(result.Close)

	return result
}
