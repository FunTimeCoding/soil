package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/generated/client"
)

type Client struct {
	context context.Context
	client  *client.Client
}
