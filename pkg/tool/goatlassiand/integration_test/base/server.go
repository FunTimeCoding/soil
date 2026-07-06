package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/mock_client"
)

type Server struct {
	MockClient    *mock_client.Client
	ContextServer *model_context_server.Server
}
