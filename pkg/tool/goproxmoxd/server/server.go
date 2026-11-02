package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

type Server struct {
	service   *service.Service
	inventory *instance.Inventory
	reporter  face.Reporter
}
