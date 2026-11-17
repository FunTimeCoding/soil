package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/openbao/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/daemon"
	"testing"
)

type Server struct {
	Source   *mock_client.Client
	Reporter *memory.Memory
	daemon   *daemon.Daemon
	served   *model_context_server.Server
	t        *testing.T
}
