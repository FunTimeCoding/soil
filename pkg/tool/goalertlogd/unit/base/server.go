package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

type Server struct {
	Daemon     *daemon.Daemon
	Clock      *worker.Worker
	MockClient *mock_client.Client
	LitePath   string
	*model_context_server.Server
}
