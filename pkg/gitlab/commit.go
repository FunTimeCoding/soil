package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Commit(
	project int64,
	branch string,
	text string,
	path string,
	content string,
	update bool,
) (*gitlab.Commit, error) {
	var action gitlab.FileActionValue

	if update {
		action = gitlab.FileUpdate
	} else {
		action = gitlab.FileCreate
	}

	result, _, e := c.client.Commits.CreateCommit(
		project,
		&gitlab.CreateCommitOptions{
			Branch:        &branch,
			CommitMessage: &text,
			Actions: []*gitlab.CommitActionOptions{
				{Action: &action, FilePath: &path, Content: &content},
			},
		},
	)

	return result, e
}
