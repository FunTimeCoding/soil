package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/worker"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	inventory *instance.Inventory
	worker    *worker.Worker
	view      *view.View
	registry  *registry.Registry
}
