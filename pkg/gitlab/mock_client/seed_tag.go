package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/tag"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SeedTag(name string) {
	c.tags = append(c.tags, tag.New(&gitlab.Tag{Name: name}))
}
