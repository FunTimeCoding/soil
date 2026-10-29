package server

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

type Server struct {
	service  *service.Service
	worker   *worker.Worker
	reporter face.Reporter
}
