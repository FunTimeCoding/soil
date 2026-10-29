package service

import (
	"github.com/funtimecoding/soil/pkg/chromium/tab"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) Wake(v *selector.Selector) (*tab.Tab, error) {
	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	if f := s.source.Wake(t.Identifier); f != nil {
		return nil, f
	}

	return t, nil
}
