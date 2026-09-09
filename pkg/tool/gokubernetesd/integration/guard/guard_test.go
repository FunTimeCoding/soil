package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/service/cluster"
	"github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func TestGuard(t *testing.T) {
	s := service.NewWithClusters(
		store.New(lite.NewMemory()),
		map[string]*cluster.Cluster{},
	)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gokubernetesd.Mount(
				s,
				true,
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyModelContext(t)
}
