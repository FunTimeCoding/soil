package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Service) baseline(
	x context.Context,
	k *watched_request.Key,
	v *snapshot.Snapshot,
) error {
	notes, e := s.notesAfter(x, k, 0)

	if e != nil {
		return e
	}

	for _, n := range notes {
		v.LastNote = max(v.LastNote, n.Identifier)
	}

	return s.store.SaveSnapshot(v)
}
