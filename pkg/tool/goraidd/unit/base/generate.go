package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/client/mock_client"
	"testing"
)

func Generate(
	t *testing.T,
	p *mock_client.Client,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()

	return Mount(t, service_tester.NewOption(t, t.TempDir(), t.TempDir(), p), r)
}
