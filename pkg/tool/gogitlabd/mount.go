package gogitlabd

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func Mount(
	c *gitlab.Client,
	b *web.Server,
	r face.Reporter,
	t face.Recorder,
	version string,
	g *guard.Mux,
) {
	model_context.New(c, r, t, version).Mount(g)
	b.Mount(g)
}
