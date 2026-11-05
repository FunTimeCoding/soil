package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/sublime/face"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/service"
)

type Daemon struct {
	client  face.SublimeClient
	guard   *peer.Guard
	service *service.Service
}
