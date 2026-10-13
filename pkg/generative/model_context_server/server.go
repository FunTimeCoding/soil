package model_context_server

import (
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Server struct {
	Port      int
	Lifecycle *lifecycle.Lifecycle
	pool      *pool.Pool
}
