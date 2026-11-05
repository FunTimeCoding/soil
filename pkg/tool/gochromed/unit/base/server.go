package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/mock_client"
)

type Server struct {
	MockClient *mock_client.Client
	Reporter   *memory.Memory
	*model_context_server.Server
}
