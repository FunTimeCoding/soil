package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
)

type Server struct {
	Daemon   *daemon.Daemon
	Forge    *mock_client.Client
	Reporter *memory.Memory
	*model_context_server.Server
}
