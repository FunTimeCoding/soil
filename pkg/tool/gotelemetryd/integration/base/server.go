package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
)

type Server struct {
	Store         *store.Store
	*model_context_server.Server
}
