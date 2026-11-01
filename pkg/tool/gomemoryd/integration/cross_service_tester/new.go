package cross_service_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/web"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/client"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/face"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/base"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Tester {
	t.Helper()
	q := base.New(t)
	q.SkipUnreachable(t)
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "memory.db")
	o.Query = func() face.Client {
		return client.New(
			locator.New(webConstant.Localhost).Insecure().Port(q.Port),
			constant.ModelContextTestToken,
		)
	}
	v := service.New(o, logger.New(t.Context()))
	r := memory.New()
	memoryServer := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gomemoryd.Mount(v, web.New(v, r), r, mock_recorder.New(), g)
		},
	)
	queryClient := q.Connect(t)
	memoryClient := memoryServer.Connect(t)
	t.Cleanup(
		func() {
			queryClient.Close()
			memoryClient.Close()
			memoryServer.Close()
			v.Close()
			q.Close()
		},
	)

	return &Tester{
		t:            t,
		QueryClient:  queryClient,
		MemoryClient: memoryClient,
		QueryServer:  q,
	}
}
