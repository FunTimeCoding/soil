package client

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
	"time"
)

type Client struct {
	generated *client.ClientWithResponses
	base      string
	token     string
	pool      *pool.Pool
	long      bool
	timeout   time.Duration
}
