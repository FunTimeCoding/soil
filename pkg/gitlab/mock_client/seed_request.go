package mock_client

import (
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"time"
)

func (c *Client) SeedRequest(
	project int64,
	identifier int64,
	state string,
	title string,
	author string,
	update time.Time,
) {
	c.requests = append(
		c.requests,
		merge_request.New(
			&gitlab.BasicMergeRequest{
				ProjectID: project,
				IID:       identifier,
				State:     state,
				Title:     title,
				Author:    &gitlab.BasicUser{Username: author},
				UpdatedAt: &update,
			},
		),
	)
}
