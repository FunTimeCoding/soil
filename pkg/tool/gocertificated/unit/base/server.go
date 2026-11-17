package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/daemon"
	"testing"
	"time"
)

type Server struct {
	Forge    *mock_client.Client
	Reporter *memory.Memory
	LitePath string
	clock    func() time.Time
	daemon   *daemon.Daemon
	served   *model_context_server.Server
	t        *testing.T
}
