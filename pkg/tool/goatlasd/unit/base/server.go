package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	outpost "github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"testing"
)

type Server struct {
	*service_tester.Tester
	Outpost *outpost.Client
	daemon  *daemon.Daemon
	served  *model_context_server.Server
	t       *testing.T
}
