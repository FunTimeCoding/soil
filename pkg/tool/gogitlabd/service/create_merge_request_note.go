package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
)

func (s *Service) CreateMergeRequestNote(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
	body string,
) (*note.Note, error) {
	return guarded(
		x,
		s,
		instance,
		func(y context.Context, c face.GitLabClient) (*note.Note, error) {
			return c.CreateMergeRequestNote(y, project, identifier, body)
		},
	)
}
