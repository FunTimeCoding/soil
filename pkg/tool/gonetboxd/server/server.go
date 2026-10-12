package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
)

type Server struct {
	client   netbox.NetboxSource
	service  *service.Service
	reporter face.Reporter
}
