package service

import (
	"github.com/funtimecoding/soil/pkg/chromium/history"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) History(v *selector.Selector) (*history.Result, error) {
	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	return s.source.Page(t.Identifier).History()
}
