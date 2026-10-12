package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/mock_client"
)

type Server struct {
	MockClient    *mock_client.Client
	Reporter      *memory.Memory
	ContextServer *model_context_server.Server
}
