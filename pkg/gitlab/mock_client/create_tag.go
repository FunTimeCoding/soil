package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) CreateTag(
	_ context.Context,
	_ int64,
	name string,
	_ string,
	_ string,
) (*tag.Tag, error) {
	result := tag.New(&gitlab.Tag{Name: name})
	c.tags = append(c.tags, result)

	return result, nil
}
