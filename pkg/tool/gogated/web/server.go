package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service       *service.Service
	authorization *client.Client
	reporter      face.Reporter
	view          *view.View
}
