package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/service"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/types/instance"
)

func New(
	v *service.Service,
	i *instance.Inventory,
	r face.Reporter,
) *Server {
	return &Server{service: v, inventory: i, reporter: r}
}
