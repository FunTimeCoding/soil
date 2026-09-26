package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/generated/client"
)

type Client struct {
	context context.Context
	client  *client.ClientWithResponses
}
