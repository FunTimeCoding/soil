package base

import (
	"github.com/funtimecoding/soil/pkg/atlassian/confluence/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
)

type Server struct {
	MockClient *mock_client.Client
	*model_context_server.Server
}
