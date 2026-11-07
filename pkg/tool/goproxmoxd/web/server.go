package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	inventory *instance.Inventory
	service   *service.Service
	view      *view.View
}
