package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"testing"
)

type Server struct {
	t      *testing.T
	daemon *daemon.Daemon
	tools  *model_context_client.Client
	*model_context_server.Server
}
