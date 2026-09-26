package server

import "github.com/funtimecoding/soil/pkg/system/service"

func New(s *service.Client) *Server {
	return &Server{service: s}
}
