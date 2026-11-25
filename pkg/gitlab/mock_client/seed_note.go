package mock_client

import (
	"gitlab.com/gitlab-org/api/client-go/v3"
	"time"
)

func (c *Client) SeedNote(
	project int64,
	request int64,
	author string,
	body string,
	system bool,
	at time.Time,
) *gitlab.Note {
	result := &gitlab.Note{
		ID:          int64(len(c.notes) + 1),
		ProjectID:   project,
		NoteableIID: request,
		Author:      gitlab.NoteAuthor{Username: author},
		Body:        body,
		System:      system,
		CreatedAt:   &at,
	}
	c.notes = append(c.notes, result)

	return result
}
