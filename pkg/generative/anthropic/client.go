package anthropic

import (
	"github.com/anthropics/anthropic-sdk-go/v2"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Client struct {
	client anthropic.Client
	pool   *pool.Pool
}
