package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"time"
)

func (s *Service) ReplyToMergeRequestDiscussion(
	x context.Context,
	instance string,
	project int64,
	identifier int64,
	thread string,
	body string,
	callsign string,
	resolve bool,
) (*note.Note, error) {
	result, e := guarded(
		x,
		s,
		instance,
		func(
			y context.Context,
			c face.GitLabClient,
		) (*note.Note, error) {
			return c.ReplyToDiscussion(y, project, identifier, thread, body)
		},
	)

	if e != nil {
		return nil, e
	}

	s.writers.Remember(result.Identifier, callsign, time.Now())

	if !resolve {
		return result, nil
	}

	_, f := guarded(
		x,
		s,
		instance,
		func(
			y context.Context,
			c face.GitLabClient,
		) (bool, error) {
			return true, c.ResolveDiscussion(y, project, identifier, thread)
		},
	)

	return result, f
}
