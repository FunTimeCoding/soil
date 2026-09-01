package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service  face.Service
	worker   *worker.Worker
	view     *view.View
	registry *palette.Registry
}
