package client

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

type Client struct {
	client *client.ClientWithResponses
	base   *locator.Locator
	token  func() string
	pool   *pool.Pool
}
