package server

import "github.com/funtimecoding/soil/pkg/tool/goprocessd/supervisor"

func New(s *supervisor.Supervisor) *Server {
	return &Server{supervisor: s}
}
