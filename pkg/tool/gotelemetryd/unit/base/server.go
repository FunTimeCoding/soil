package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Server struct {
	Store    *store.Store
	Daemon   *daemon.Daemon
	Reporter *memory.Memory
	*model_context_server.Server
}
