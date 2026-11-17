package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/source/inventory"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/daemon"
	"testing"
)

type Server struct {
	Memory         *mock_client.Client
	Reporter       *memory.Memory
	Directory      string
	inventory      *inventory.Inventory
	indexDirectory string
	daemon         *daemon.Daemon
	served         *model_context_server.Server
	t              *testing.T
}
