package daemon

import (
	"github.com/funtimecoding/soil/pkg/github/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/worker"
)

type Daemon struct {
	client  face.GitHubClient
	guard   *peer.Guard
	service *service.Service
	clock   *worker.Worker
}
