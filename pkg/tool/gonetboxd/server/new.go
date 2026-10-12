package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/service"
)

func New(
	c netbox.NetboxSource,
	v *service.Service,
	r face.Reporter,
) *Server {
	return &Server{client: c, service: v, reporter: r}
}
