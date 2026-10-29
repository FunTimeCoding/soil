package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListChoiceSets(_ context.Context) ([]*client.ChoiceSet, error) {
	return nil, c.Failure
}
