package client

import (
	"github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Client struct {
	client *client.ClientWithResponses
	base   *locator.Locator
	token  string
	pool   *pool.Pool
}
