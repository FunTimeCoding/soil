package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
)

type Server struct {
	service   *service.Service
	inventory *instance.Inventory
	reporter  face.Reporter
}
