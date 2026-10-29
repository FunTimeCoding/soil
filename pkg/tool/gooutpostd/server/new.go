package server

import "github.com/funtimecoding/soil/pkg/tool/gooutpostd/service"

func New(v *service.Service) *Server {
	return &Server{service: v}
}
