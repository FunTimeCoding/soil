package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/fill"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"
)

func (s *Service) Fill(r *fill.Fill) (*view.View, error) {
	if r.UID == "" {
		return nil, validation.New("uid is required")
	}

	if r.Value == "" {
		return nil, validation.New("value is required")
	}

	t, e := s.resolveTab(r.Tab)

	if e != nil {
		return nil, e
	}

	if f := s.enter(t, r); f != nil {
		return nil, f
	}

	if !r.Snapshot {
		return view.New(""), nil
	}

	return s.observe(t)
}
