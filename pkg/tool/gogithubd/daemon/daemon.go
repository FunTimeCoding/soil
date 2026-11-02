package daemon

import (
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/face"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
)

type Daemon struct {
	source  face.RunSource
	service *service.Service
	clock   *worker.Worker
}
