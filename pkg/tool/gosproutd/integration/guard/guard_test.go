package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/integration/mock_notifier"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd"
	sproutConstant "github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
	"time"
)

func TestGuard(t *testing.T) {
	now := time.Now().UTC()
	v := service.New(
		store.New(lite.NewMemory(), func() time.Time { return now }),
		mock_notifier.New(),
	)
	w := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gosproutd.Mount(
				v,
				web.New(v),
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)
	defer w.Stop()
	w.VerifyBase(t)
	w.VerifyOpen(t, sproutConstant.DashboardPath)
	w.VerifyModelContext(t)
}
