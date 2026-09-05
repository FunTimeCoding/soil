package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
)

type Server struct {
	Store  *store.Store
	server *model_context_server.Server
}
