package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web"
)

type Daemon struct {
	guard   *peer.Guard
	service *service.Service
	web     *web.Server
}
