package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
)

func (c *Client) Tree(
	_ context.Context,
	_ int64,
	_ string,
	_ string,
	_ bool,
	_ int64,
) ([]*tree.Node, error) {
	return nil, nil
}
