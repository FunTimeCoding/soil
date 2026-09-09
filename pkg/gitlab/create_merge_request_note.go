package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) CreateMergeRequestNote(
	project int64,
	identifier int64,
	body string,
) (*note.Note, error) {
	result, _, e := c.client.Notes.CreateMergeRequestNote(
		project,
		identifier,
		&gitlab.CreateMergeRequestNoteOptions{Body: &body},
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return note.New(result), nil
}
