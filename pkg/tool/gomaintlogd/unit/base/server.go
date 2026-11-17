package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"testing"
)

type Server struct {
	Reporter *memory.Memory
	LitePath string
	t        *testing.T
	*model_context_server.Server
}
