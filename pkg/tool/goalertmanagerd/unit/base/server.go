package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"testing"
)

type Server struct {
	*service_tester.Tester
	daemon *daemon.Daemon
	served *model_context_server.Server
	t      *testing.T
}
