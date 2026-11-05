package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/iterm/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goitermd/daemon"
)

type Server struct {
	MockClient *mock_client.Client
	Reporter   *memory.Memory
	Daemon     *daemon.Daemon
	*model_context_server.Server
}
