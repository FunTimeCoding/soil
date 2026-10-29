package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) GetChoiceSet(_ context.Context, _ string) (client.ChoiceSet, error) {
	return client.ChoiceSet{}, c.Failure
}
