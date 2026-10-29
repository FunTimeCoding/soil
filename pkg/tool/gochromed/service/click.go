package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/view"
)

func (s *Service) Click(
	v *selector.Selector,
	uid string,
	snapshot bool,
) (*view.View, error) {
	if uid == "" {
		return nil, validation.New("uid is required")
	}

	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	node, f := s.resolveUID(t.Identifier, uid)

	if f != nil {
		return nil, f
	}

	p := s.source.Page(t.Identifier)

	if g := p.ClickNode(node); g != nil {
		return nil, g
	}

	if !snapshot {
		return view.New(nil, nil), nil
	}

	return s.observe(t.Identifier, p)
}
