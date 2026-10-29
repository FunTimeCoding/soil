package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/service"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
)

type Server struct {
	Store    *store.Store
	Service  *service.Service
	System   *mock_system.System
	Reporter *memory.Memory
	Client   *client.ClientWithResponses
	*model_context_server.Server
}
