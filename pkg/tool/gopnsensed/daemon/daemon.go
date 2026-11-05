package daemon

import (
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/service"
)

type Daemon struct {
	client  face.OpnsenseClient
	guard   *peer.Guard
	service *service.Service
}
