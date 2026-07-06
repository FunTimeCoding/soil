package model_context_server

import "github.com/funtimecoding/soil/pkg/lifecycle"

type Server struct {
	Port      int
	Lifecycle *lifecycle.Lifecycle
}
