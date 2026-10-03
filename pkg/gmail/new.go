package gmail

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/authorization/callback"
)

func New(directory string) *Client {
	return &Client{
		context:   context.Background(),
		directory: directory,
		callback:  callback.NewEnvironment(false),
	}
}
