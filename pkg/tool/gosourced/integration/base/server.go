package base

import "github.com/funtimecoding/soil/pkg/generative/model_context_server"

type Server struct {
	*model_context_server.Server
	Directory     string
}
