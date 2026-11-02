package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/unit/mock_system"
	"gorm.io/gorm"
)

type Server struct {
	Mapper   *gorm.DB
	System   *mock_system.System
	Reporter *memory.Memory
	Client   *client.ClientWithResponses
	*model_context_server.Server
}
