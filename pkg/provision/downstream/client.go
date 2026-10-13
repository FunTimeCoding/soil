package downstream

import "github.com/funtimecoding/soil/pkg/web/stall_client/pool"

type Client struct {
	host string
	port int
	pool *pool.Pool
}
