package daemon

import (
	"github.com/funtimecoding/soil/pkg/linkace/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/service"
)

type Daemon struct {
	client  face.LinkAceClient
	guard   *peer.Guard
	service *service.Service
}
