package web

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/soil/pkg/web/palette"
	"github.com/funtimecoding/soil/pkg/web/view"
)

type Server struct {
	store    *store.Store
	notifier face.EventNotifier
	worker   *worker.Worker
	view     *view.View
	registry *palette.Registry
}
