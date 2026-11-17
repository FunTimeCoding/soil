package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"testing"
)

type Server struct {
	Daemon       *daemon.Daemon
	Clock        *worker.Worker
	Alertmanager *mock_client.Client
	LitePath     string
	t            *testing.T
	*model_context_server.Server
}
