package n8n

import "github.com/funtimecoding/soil/pkg/web/stall_client/pool"

type Client struct {
	host  string
	token string
	pool  *pool.Pool
}
