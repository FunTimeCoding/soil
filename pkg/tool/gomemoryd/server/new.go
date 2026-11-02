package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
)

func New(
	s *service.Service,
	hiddenTag string,
	r face.Reporter,
) *Server {
	return &Server{service: s, hiddenTag: hiddenTag, reporter: r}
}
