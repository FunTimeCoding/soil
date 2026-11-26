package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goladybirdd/types/view"
)

func (s *Service) Click(
	tab string,
	uid string,
	snapshot bool,
) (*view.View, error) {
	if uid == "" {
		return nil, validation.New("uid is required")
	}

	t, e := s.resolveTab(tab)

	if e != nil {
		return nil, e
	}

	mark := s.navigation.Mark(t)

	if f := s.act(t, uid, constant.ClickScript); f != nil {
		return nil, f
	}

	result := view.New("")

	if s.navigation.StartedSince(t, mark) {
		result.Loading = !s.navigation.Settle(t)
		result.Landed = s.navigation.Locator(t)
	}

	if !snapshot {
		return result, nil
	}

	observed, g := s.observe(t)

	if g != nil {
		return nil, g
	}

	observed.Landed = result.Landed

	return observed, nil
}
