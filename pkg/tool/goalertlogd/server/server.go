package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

type Server struct {
	store    *store.Store
	worker   *worker.Worker
	reporter face.Reporter
}
