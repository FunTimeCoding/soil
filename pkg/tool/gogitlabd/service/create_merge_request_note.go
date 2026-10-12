package service

import "github.com/funtimecoding/soil/pkg/gitlab/note"

func (s *Service) CreateMergeRequestNote(
	project int64,
	identifier int64,
	body string,
) (*note.Note, error) {
	return s.client.CreateMergeRequestNote(project, identifier, body)
}
