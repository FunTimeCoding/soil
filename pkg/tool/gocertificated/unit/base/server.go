package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
)

type Server struct {
	Service  *service.Service
	Forge    *mock_client.Client
	Reporter *memory.Memory
	*model_context_server.Server
}
