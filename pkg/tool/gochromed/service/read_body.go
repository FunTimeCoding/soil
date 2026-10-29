package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/saved"
	"github.com/funtimecoding/soil/pkg/tool/gochromed/types/selector"
)

func (s *Service) ReadBody(v *selector.Selector) (*saved.Saved, error) {
	t, e := s.resolveTab(v)

	if e != nil {
		return nil, e
	}

	body, f := s.source.Page(t.Identifier).Body()

	if f != nil {
		return nil, f
	}

	return s.save(t, "html", []byte(body))
}
