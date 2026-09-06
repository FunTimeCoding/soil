package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/mock_client"
)

type Server struct {
	MockClient    *mock_client.Client
	*model_context_server.Server
}
