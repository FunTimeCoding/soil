package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) MergeRequestNotes(
	_ context.Context,
	project int64,
	identifier int64,
	after int64,
) ([]*note.Note, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	var result []*gitlab.Note

	for _, v := range c.notes {
		if v.ProjectID == project && v.NoteableIID == identifier &&
			v.ID > after {
			result = append(result, v)
		}
	}

	return note.NewSlice(result), nil
}
