package mock_client

import "github.com/funtimecoding/soil/pkg/linkace/tag"

func (c *Client) UpdateTag(
	_ int,
	_ map[string]any,
) (*tag.Tag, error) {
	return nil, nil
}
