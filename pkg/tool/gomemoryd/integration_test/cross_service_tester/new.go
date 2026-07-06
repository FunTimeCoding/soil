package cross_service_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
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
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	q := base.New(t)
	c, e := client.NewClient(
		fmt.Sprintf("http://localhost:%d", q.Port()),
	)
	errors.PanicOnError(e)
	idx := memory_indexer.New(c)
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	v := service.New(s, idx, idx, idx)
	r := memory.New()
	gomemorydServer := model_context_server.New(
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
	t.Cleanup(
		func() {
			gomemorydServer.Stop()
			s.Close()
			q.Close()
		},
	)

	return &Tester{
		Goqueryd:       model_context_client.New(t, q.Port()),
		Gomemoryd:      model_context_client.New(t, gomemorydServer.Port),
		GomemorydStore: s,
		goqueryd:       q,
	}
}
