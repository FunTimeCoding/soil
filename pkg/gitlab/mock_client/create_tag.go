package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CreateTag(
	_ int64,
	name string,
	_ string,
	_ string,
) (*tag.Tag, error) {
	return tag.New(&gitlab.Tag{Name: name}), nil
}
