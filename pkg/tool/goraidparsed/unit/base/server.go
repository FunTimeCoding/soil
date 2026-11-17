package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/option"
	"testing"
)

type Server struct {
	Reporter *memory.Memory
	option   *option.Parser
	served   *model_context_server.Server
	t        *testing.T
}
