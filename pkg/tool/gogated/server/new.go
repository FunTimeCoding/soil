package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(
	s *service.Service,
	r face.Reporter,
) *Server {
	return &Server{
		service:  s,
		reporter: r,
		view:     view.New(layout.New(constant.Identity)),
	}
}
