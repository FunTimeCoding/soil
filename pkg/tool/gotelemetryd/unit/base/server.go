package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"testing"
	"time"
)

type Server struct {
	Reporter *memory.Memory
	clock    func() time.Time
	litePath string
	daemon   *daemon.Daemon
	store    *store.Store
	served   *model_context_server.Server
	t        *testing.T
}
