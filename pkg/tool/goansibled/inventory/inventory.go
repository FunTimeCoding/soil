package inventory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Inventory struct {
	context context.Context
	client  *client.ClientWithResponses
	pool    *pool.Pool
}
