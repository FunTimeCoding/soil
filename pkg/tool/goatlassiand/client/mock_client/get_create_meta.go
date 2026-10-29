package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetCreateMeta(
	_ context.Context,
	_ string,
	_ string,
	_ string,
) ([]*client.CreateMetaField, error) {
	return nil, c.Failure
}
