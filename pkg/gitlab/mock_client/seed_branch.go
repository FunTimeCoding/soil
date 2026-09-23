package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/branch"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"time"
)

func (c *Client) SeedBranch(
	name string,
	hash string,
) {
	c.branches = append(
		c.branches,
		branch.New(
			&gitlab.Branch{
				Name:   name,
				Commit: &gitlab.Commit{ID: hash, CreatedAt: new(time.Now())},
			},
		),
	)
}
