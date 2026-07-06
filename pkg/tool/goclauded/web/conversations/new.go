package conversations

import "github.com/funtimecoding/soil/pkg/tool/goclauded/service"

func New(s *service.Service) *Server {
	return &Server{service: s}
}
