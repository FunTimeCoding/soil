package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) CreateMergeRequestNote(
	project int64,
	mergeRequest int64,
	body string,
) (*gitlab.Note, error) {
	result, _, e := c.client.Notes.CreateMergeRequestNote(
		project,
		mergeRequest,
		&gitlab.CreateMergeRequestNoteOptions{Body: &body},
	)

	return result, wrapError(e)
}
