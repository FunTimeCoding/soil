package web

import (
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/service"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service *service.Service
	view    *view.View
}
