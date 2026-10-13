package sublime

import (
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"net/http"
)

type Client struct {
	base   *locator.Locator
	client *http.Client
	pool   *pool.Pool
}
