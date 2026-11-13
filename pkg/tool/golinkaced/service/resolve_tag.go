package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/linkace/tag"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) resolveTag(name string) (int, error) {
	t, e := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*tag.Tag, error) { return s.client.Tag(name) },
	)

	if e != nil {
		return 0, e
	}

	return t.Identifier, nil
}
