package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/service"
)

type Server struct {
	Service  *service.Service
	Reporter *memory.Memory
	*model_context_server.Server
}
