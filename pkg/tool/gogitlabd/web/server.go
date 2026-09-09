package web

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	client   face.Forge
	worker   *worker.Worker
	view     *view.View
	registry *palette.Registry
}
