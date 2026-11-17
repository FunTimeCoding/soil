package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gooutpostd/client/mock_client"
	"testing"
)

type Server struct {
	*service_tester.Tester
	Outpost *mock_client.Client
	daemon  *daemon.Daemon
	served  *model_context_server.Server
	t       *testing.T
}
