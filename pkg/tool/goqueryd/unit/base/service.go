package base

import "github.com/funtimecoding/soil/pkg/tool/goqueryd/service"

func (s *Server) Service() *service.Service {
	return s.service
}
