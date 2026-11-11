package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/service"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/worker"
)

type Daemon struct {
	guards  []*peer.Guard
	service *service.Service
	clock   *worker.Worker
}
