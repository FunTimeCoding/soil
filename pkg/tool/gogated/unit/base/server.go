package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
)

type Server struct {
	Store   *store.Store
	Service *service.Service
	Web     *model_context_server.Server
}
