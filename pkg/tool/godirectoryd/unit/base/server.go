package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
)

type Server struct {
	Reporter *memory.Memory
	*model_context_server.Server
}
