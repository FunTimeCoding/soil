package pushover

import "github.com/funtimecoding/soil/pkg/web/stall_client/pool"

type Client struct {
	user  string
	token string
	pool  *pool.Pool
}
