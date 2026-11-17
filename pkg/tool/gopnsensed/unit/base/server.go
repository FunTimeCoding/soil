package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/opnsense/face"
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/daemon"
	"testing"
)

type Server struct {
	Reporter *memory.Memory
	source   face.OpnsenseClient
	daemon   *daemon.Daemon
	served   *model_context_server.Server
	t        *testing.T
}
