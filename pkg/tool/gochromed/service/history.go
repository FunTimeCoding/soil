package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) History(v *selector.Selector) (*history.Result, error) {
	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	return peer.Guarded(
		context.Background(),
		s.guard,
		func(context.Context) (*history.Result, error) {
			return s.source.Page(t.Identifier).History()
		},
	)
}
