package cross_service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/memory_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/integration_test/base"
	"net/http"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	q := base.New(t)
	c, e := client.NewClient(
		fmt.Sprintf("http://localhost:%d", q.Port()),
	)
	errors.PanicOnError(e)
	i := memory_indexer.New(c)
	s := store.New(connection.NewMemory())
	v := service.New(s, i, i, i)
	r := memory.New()
	memoryServer := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			generated.HandlerFromMux(
				generated.NewStrictHandler(server.New(v, r), nil),
				m,
			)
			model_context.New(
				v,
				r,
				mock_recorder.New(),
				constant.DefaultVersion,
			).Mount(m)
		},
	)
	queryClient := model_context_client.New(t, q.Port())
	memoryClient := model_context_client.New(t, memoryServer.Port)
	t.Cleanup(
		func() {
			queryClient.Close()
			memoryClient.Close()
			memoryServer.Stop()
			s.Close()
			q.Close()
		},
	)

	return &Tester{
		QueryClient:  queryClient,
		MemoryClient: memoryClient,
		MemoryStore:  s,
		QueryServer:  q,
	}
}
