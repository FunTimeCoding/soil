package note

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(v *gitlab.Note) *Note {
	return &Note{
		Identifier: v.ID,
		Author:     v.Author.Username,
		Body:       v.Body,
		System:     v.System,
		Create:     v.CreatedAt,
		Raw:        v,
	}
}
