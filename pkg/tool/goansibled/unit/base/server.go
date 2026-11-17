package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/unit/ansible_tester"
	"testing"
)

type Server struct {
	*ansible_tester.Tester
	daemon *daemon.Daemon
	served *model_context_server.Server
	t      *testing.T
}
