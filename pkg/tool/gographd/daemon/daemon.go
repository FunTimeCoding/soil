package daemon

import (
	"github.com/funtimecoding/soil/pkg/chat/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gographd/service"
)

type Daemon struct {
	client  face.TeamsClient
	guard   *peer.Guard
	service *service.Service
}
