package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/service"
)

type Server struct {
	MockClient *mock_client.Client
	Service    *service.Service
	Reporter   *memory.Memory
	Directory  string
	*model_context_server.Server
}
