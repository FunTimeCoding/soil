package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/saved"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) Screenshot(v *selector.Selector) (*saved.Saved, error) {
	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	b, f := peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) ([]byte, error) {
			return s.source.Page(t.Identifier).Screenshot()
		},
	)

	if f != nil {
		return nil, f
	}

	return s.save(t, "png", b)
}
