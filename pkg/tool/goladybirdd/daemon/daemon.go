package daemon

import (
	"github.com/funtimecoding/soil/pkg/browser/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/service"
)

type Daemon struct {
	client  face.LadybirdClient
	guard   *peer.Guard
	service *service.Service
}
