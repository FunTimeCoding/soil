package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	store         *store.Store
	service       *service.Service
	authorization *client.Client
	view          *view.View
	registry      *palette.Registry
}
