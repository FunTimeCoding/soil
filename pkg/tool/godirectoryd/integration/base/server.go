package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/service"
	"github.com/funtimecoding/soil/pkg/web/authorization/client"
)

type Server struct {
	Service       *service.Service
	Authorization *client.Client
	*model_context_server.Server
}
