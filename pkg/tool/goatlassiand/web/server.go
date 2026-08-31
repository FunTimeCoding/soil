package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/worker"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	worker   *worker.Worker
	view     *view.View
	registry *palette.Registry
}
