package author

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/studio-b12/gowebdav"
)

type Client struct {
	client *gowebdav.Client
	pool   *pool.Pool
}
