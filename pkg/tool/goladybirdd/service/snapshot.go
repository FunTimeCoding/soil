package service

import "github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"

func (s *Service) Snapshot(tab string) (*view.View, error) {
	t, e := s.resolveTab(tab)

	if e != nil {
		return nil, e
	}

	return s.observe(t)
}
