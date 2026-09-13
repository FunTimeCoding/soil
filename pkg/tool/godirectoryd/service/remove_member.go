package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
	"slices"
)

func (s *Service) RemoveMember(
	name string,
	account string,
) (*group.Group, error) {
	current, e := s.Group(name)

	if e != nil {
		return nil, e
	}

	if !slices.Contains(current.Member, account) {
		return nil, not_found.New(constant.MemberAttribute, account)
	}

	remaining := slices.DeleteFunc(
		slices.Clone(current.Member),
		func(m string) bool { return m == account },
	)
	f := s.directory.Modify(
		s.groupName(name),
		map[string][]string{constant.MemberAttribute: remaining},
	)

	if f != nil {
		return nil, f
	}

	return s.Group(name)
}
