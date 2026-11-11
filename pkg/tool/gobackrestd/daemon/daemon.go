package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/service"
	"github.com/funtimecoding/soil/pkg/tool/gobackrestd/worker"
)

type Daemon struct {
	guard   *peer.Guard
	service *service.Service
	clock   *worker.Worker
}
