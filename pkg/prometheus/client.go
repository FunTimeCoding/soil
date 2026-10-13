package prometheus

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/prometheus/client_golang/api/prometheus/v1"
)

type Client struct {
	context            context.Context
	client             v1.API
	host               string
	alternateGraphHost string
	pool               *pool.Pool
}
