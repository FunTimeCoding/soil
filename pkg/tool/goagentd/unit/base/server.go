package base

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/daemon"
	"io"
	"testing"
)

type Server struct {
	stream func(string, io.Writer, ...string) face.Stream
	daemon *daemon.Daemon
	served *model_context_server.Server
	t      *testing.T
}
