package mock_client

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CommitActions(
	_ int64,
	branch string,
	message string,
	v []*gitlab.CommitActionOptions,
) (*gitlab.Commit, error) {
	c.commits = append(
		c.commits,
		&Commit{Branch: branch, Message: message, Actions: v},
	)

	return &gitlab.Commit{}, nil
}
