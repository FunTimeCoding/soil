package daemon

import (
	"github.com/funtimecoding/soil/pkg/habitica/face"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/service"
)

type Daemon struct {
	client  face.HabiticaClient
	guard   *peer.Guard
	service *service.Service
}
