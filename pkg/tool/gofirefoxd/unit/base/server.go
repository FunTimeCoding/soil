package base

import (
	"github.com/funtimecoding/soil/pkg/browser/firefox/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gofirefoxd/daemon"
	"testing"
)

type Server struct {
	Source    *mock_client.Client
	Daemon    *daemon.Daemon
	Reporter  *memory.Memory
	Directory string
	t         *testing.T
	*model_context_server.Server
}
