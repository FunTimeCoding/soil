package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

type Client struct {
	context  context.Context
	client   *client.ClientWithResponses
	instance string
}
