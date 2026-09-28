package server

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/layout"
	"github.com/funtimecoding/soil/pkg/web/view"
)

func New(s *service.Service) *Server {
	return &Server{service: s, view: view.New(layout.New(constant.Identity))}
}
