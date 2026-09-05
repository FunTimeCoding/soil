package base

import "github.com/funtimecoding/soil/pkg/generative/model_context_server"

type Server struct {
	ContextServer *model_context_server.Server
	Directory     string
}
