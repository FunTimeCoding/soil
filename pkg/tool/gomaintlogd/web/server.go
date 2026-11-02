package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
	"github.com/funtimecoding/soil/pkg/web/palette/registry"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	service  *service.Service
	notifier face.EventNotifier
	reporter face.Reporter
	view     *view.View
	registry *registry.Registry
}
