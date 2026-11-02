package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service       *service.Service
	authorization *client.Client
	registry      *registry.Registry
	view          *view.View
}
