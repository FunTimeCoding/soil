package assistant

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assistant/connection"
)

func New(
	host string,
	token string,
	o ...Option,
) *Client {
	result := &Client{
		connection: connection.New(host, token),
		context:    context.Background(),
	}

	for _, o := range o {
		o(result)
	}

	return result
}
