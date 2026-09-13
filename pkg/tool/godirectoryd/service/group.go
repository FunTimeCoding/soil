package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
)

func (s *Service) Group(name string) (*group.Group, error) {
	found, e := s.Groups()

	if e != nil {
		return nil, e
	}

	for _, g := range found {
		if g.Name == name {
			return g, nil
		}
	}

	return nil, not_found.New(constant.GroupKind, name)
}
