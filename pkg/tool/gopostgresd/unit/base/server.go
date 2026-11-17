package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/unit/service_tester"
	"testing"
)

type Server struct {
	*service_tester.Tester
	BackupDirectory string
	daemon          *daemon.Daemon
	served          *model_context_server.Server
	t               *testing.T
}
