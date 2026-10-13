package hetzner

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

type Client struct {
	context context.Context
	client  *hcloud.Client
	pool    *pool.Pool
}
