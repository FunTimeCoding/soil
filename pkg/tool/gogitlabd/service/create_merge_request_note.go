package service

import (
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
)

func (s *Service) CreateMergeRequestNote(
	instance string,
	project int64,
	identifier int64,
	body string,
) (*note.Note, error) {
	return guarded(
		s,
		instance,
		func(c face.GitLabClient) (*note.Note, error) {
			return c.CreateMergeRequestNote(project, identifier, body)
		},
	)
}
