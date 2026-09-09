package mock_client

import "github.com/funtimecoding/soil/pkg/gitlab/tag"

func (c *Client) MustTags(_ int64) []*tag.Tag {
	return nil
}
