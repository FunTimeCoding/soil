package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service       *service.Service
	authorization *client.Client
	superUserMail string
	view          *view.View
	registry      *palette.Registry
}
