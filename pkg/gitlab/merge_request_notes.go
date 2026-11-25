package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"gitlab.com/gitlab-org/api/client-go/v3"
	"slices"
)

func (c *Client) MergeRequestNotes(
	x context.Context,
	project int64,
	identifier int64,
	after int64,
) ([]*note.Note, error) {
	var result []*gitlab.Note

	for page := range int64(constant.NotePageLimit) {
		v, _, e := c.client.Notes.ListMergeRequestNotes(
			project,
			identifier,
			&gitlab.ListMergeRequestNotesOptions{
				ListOptions: gitlab.ListOptions{
					Page:    page + 1,
					PerPage: constant.NotePageSize,
				},
				OrderBy: new(constant.CreatedOrder),
				Sort:    new(constant.Descending),
			},
			gitlab.WithContext(x),
		)

		if e != nil {
			return nil, wrapError(e)
		}

		done := len(v) < constant.NotePageSize

		for _, n := range v {
			if n.ID <= after {
				done = true

				break
			}

			result = append(result, n)
		}

		if done {
			break
		}
	}

	slices.Reverse(result)

	return note.NewSlice(result), nil
}
