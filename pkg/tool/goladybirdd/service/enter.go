package service

import (
	"github.com/funtimecoding/soil/pkg/browser/types/request"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/fill"
)

func (s *Service) enter(
	tab string,
	r *fill.Fill,
) error {
	if r.Direct {
		return s.act(
			tab,
			r.UID,
			constant.DirectScript,
			request.NewText(r.Value),
		)
	}

	if e := s.act(tab, r.UID, constant.FocusScript); e != nil {
		return e
	}

	return s.attempt(
		func() error {
			return s.client.Type(tab, r.Value)
		},
	)
}
