package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service *service.Service
	view    *view.View
}
