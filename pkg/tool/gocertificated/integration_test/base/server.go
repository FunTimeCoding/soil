package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/store"
)

type Server struct {
	Store   *store.Store
	Service *service.Service
	Forge   *mock_client.Client
	server  *model_context_server.Server
}
