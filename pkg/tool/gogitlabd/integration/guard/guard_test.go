package guard

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	soilGitlab "github.com/funtimecoding/soil/pkg/gitlab"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd"
	gitlabConstant "github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"net/http"
	"testing"
	"time"
)

func TestGuard(t *testing.T) {
	var c *soilGitlab.Client
	var n *gitlab.Client
	k := worker.New(
		c,
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		memory.New(),
	)
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			gogitlabd.Mount(
				n,
				web.New(c, k),
				memory.New(),
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)
	defer v.Stop()
	v.VerifyBase(t)
	v.VerifyOpen(t, gitlabConstant.BoardPath)
	v.VerifyOpen(t, "/event")
	v.VerifyModelContext(t)
}
