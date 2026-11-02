package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/inventory/instance"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/service"
)

func New(
	s *service.Service,
	i *instance.Inventory,
	r face.Reporter,
) *Server {
	return &Server{service: s, inventory: i, reporter: r}
}
