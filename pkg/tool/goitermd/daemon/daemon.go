package daemon

import (
	"github.com/funtimecoding/soil/pkg/iterm/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/service"
)

type Daemon struct {
	client  face.ItermClient
	guard   *peer.Guard
	service *service.Service
}
