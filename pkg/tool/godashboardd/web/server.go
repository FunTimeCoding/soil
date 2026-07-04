package web

import (
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/go-library/pkg/tool/godashboardd/store"
	"github.com/funtimecoding/go-library/pkg/web/authorization/client"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	"github.com/funtimecoding/go-library/pkg/web/view"
)

type Server struct {
	board         *board.Board
	service       *service.Service
	store         *store.Store
	authorization *client.Client
	labels        map[string]bool
	view          *view.View
	registry      *palette.Registry
}
