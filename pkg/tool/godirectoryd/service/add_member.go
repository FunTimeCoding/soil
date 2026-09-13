package service

import (
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
	"slices"
)

func (s *Service) AddMember(
	name string,
	account string,
) (*group.Group, error) {
	current, e := s.Group(name)

	if e != nil {
		return nil, e
	}

	if _, f := s.User(account); f != nil {
		return nil, f
	}

	if slices.Contains(current.Member, account) {
		return nil, conflict.Exists(constant.MemberAttribute, account)
	}

	g := s.directory.Modify(
		s.groupName(name),
		map[string][]string{
			constant.MemberAttribute: append(current.Member, account),
		},
	)

	if g != nil {
		return nil, g
	}

	return s.Group(name)
}
