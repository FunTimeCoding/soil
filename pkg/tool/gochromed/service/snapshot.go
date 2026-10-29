package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

func (s *Service) Snapshot(v *selector.Selector) (*view.View, error) {
	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	return s.observe(t.Identifier, s.source.Page(t.Identifier))
}
