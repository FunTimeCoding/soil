package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/mock_client"
)

type Server struct {
	Alfa   *mock_client.Client
	Bravo  *mock_client.Client
	Client *client.ClientWithResponses
	*model_context_server.Server
}
