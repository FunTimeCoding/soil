package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service   *service.Service
	inventory *instance.Inventory
	reporter  face.Reporter
	view      *view.View
}
