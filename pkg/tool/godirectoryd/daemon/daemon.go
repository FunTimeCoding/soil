package daemon

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/web"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
)

type Daemon struct {
	guard         *peer.Guard
	authorization *client.Client
	service       *service.Service
	web           *web.Server
}
