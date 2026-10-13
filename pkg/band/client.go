package band

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"net/http"
)

type Client struct {
	base     string
	user     string
	password string
	client   *http.Client
	pool     *pool.Pool
}
