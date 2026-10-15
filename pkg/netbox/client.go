package netbox

import (
	"context"
	"github.com/funtimecoding/soil/pkg/netbox/cache"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/netbox-community/go-netbox/v4"
	"github.com/prometheus/client_golang/prometheus"
)

type Client struct {
	context context.Context
	client  *netbox.APIClient
	cache   *cache.Cache
	verbose  bool
	observer prometheus.ObserverVec
	pool     *pool.Pool
}
