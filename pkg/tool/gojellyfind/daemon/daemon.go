package daemon

import (
	"github.com/funtimecoding/soil/pkg/jellyfin/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/service"
)

type Daemon struct {
	client  face.JellyfinClient
	guard   *peer.Guard
	service *service.Service
}
