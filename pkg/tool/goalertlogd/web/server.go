package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service  *service.Service
	notifier face.EventNotifier
	view     *view.View
}
