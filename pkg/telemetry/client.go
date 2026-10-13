package telemetry

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"net/http"
	"sync"
)

type Client struct {
	base   string
	client *http.Client
	pool   *pool.Pool
	group  sync.WaitGroup
}
