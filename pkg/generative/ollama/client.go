package ollama

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/ollama/ollama/api"
)

type Client struct {
	context context.Context
	client  *api.Client
	host    string
	port    int
	secure  bool
	model   string
	pool    *pool.Pool
}
