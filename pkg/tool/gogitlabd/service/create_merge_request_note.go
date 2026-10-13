package service

import "github.com/funtimecoding/soil/pkg/gitlab/note"

func (s *Service) CreateMergeRequestNote(
	instance string,
	project int64,
	identifier int64,
	body string,
) (*note.Note, error) {
	c, e := s.pool.Client(instance)

	if e != nil {
		return nil, e
	}

	return c.CreateMergeRequestNote(project, identifier, body)
}
