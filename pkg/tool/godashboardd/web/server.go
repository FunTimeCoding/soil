package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/service"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	board    *board.Board
	service  *service.Service
	notifier face.EventNotifier
	reporter face.Reporter
	view     *view.View
}
