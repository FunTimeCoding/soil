package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"time"
)

func (s *Service) CreateMergeRequestNote(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
	body string,
	callsign string,
) (*note.Note, error) {
	result, e := guarded(
		x,
		s,
		instance,
		func(
			y context.Context,
			c face.GitLabClient,
		) (*note.Note, error) {
			return c.CreateMergeRequestNote(y, project, identifier, body)
		},
	)

	if e != nil {
		return nil, e
	}

	s.writers.Remember(result.Identifier, callsign, time.Now())

	return result, nil
}
