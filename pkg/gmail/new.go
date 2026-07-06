package gmail

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/authorization/callback"
)

func New(directory string) *Client {
	// TODO: If port is busy, fail or block until it is free
	return &Client{
		context:   context.Background(),
		directory: directory,
		callback:  callback.NewEnvironment(false),
	}
}
