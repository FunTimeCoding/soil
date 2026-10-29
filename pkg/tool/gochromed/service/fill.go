package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/fill"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

func (s *Service) Fill(r *fill.Fill) (*view.View, error) {
	if r.UID == "" {
		return nil, validation.New("uid is required")
	}

	if r.Value == "" {
		return nil, validation.New("value is required")
	}

	t, e := s.resolveTab(r.Selector)

	if e != nil {
		return nil, e
	}

	node, f := s.resolveUID(t.Identifier, r.UID)

	if f != nil {
		return nil, f
	}

	p := s.source.Page(t.Identifier)

	if g := p.FillNode(node, r.Value, r.Direct); g != nil {
		return nil, g
	}

	if !r.Snapshot {
		return view.New(nil, nil), nil
	}

	return s.observe(t.Identifier, p)
}
