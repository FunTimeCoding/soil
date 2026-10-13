package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
	pool    *pool.Pool
}
