package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store"
	"testing"
)

type Server struct {
	Reporter      *memory.Memory
	litePath      string
	seedDirectory string
	daemon        *daemon.Daemon
	served        *model_context_server.Server
	store         *store.Store
	t             *testing.T
}
