package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/face"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Service) notesAfter(
	x context.Context,
	k *watched_request.Key,
	after int64,
) ([]*note.Note, error) {
	return guarded(
		x,
		s,
		k.Instance,
		func(
			y context.Context,
			c face.GitLabClient,
		) ([]*note.Note, error) {
			return c.MergeRequestNotes(y, k.Project, k.Request, after)
		},
	)
}
