package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) Navigate(
	v *selector.Selector,
	l string,
) error {
	if l == "" {
		return validation.New("url is required")
	}

	t, e := s.resolveTab(v)

	if e != nil {
		return e
	}

	return s.source.Page(t.Identifier).Navigate(l)
}
