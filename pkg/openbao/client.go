package openbao

import (
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"github.com/openbao/openbao/api/v2"
	"sync"
	"time"
)

type Client struct {
	client    *api.Client
	role      string
	tokenPath string
	expiry    time.Time
	lock      sync.Mutex
	pool      *pool.Pool
}
