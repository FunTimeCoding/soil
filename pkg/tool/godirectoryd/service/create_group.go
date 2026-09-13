package service

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
	"strconv"
)

func (s *Service) CreateGroup(name string) (*group.Group, error) {
	number, e := s.nextGroupNumber()

	if e != nil {
		return nil, e
	}

	f := s.directory.Add(
		s.groupName(name),
		map[string][]string{
			constant.ClassAttribute:       {constant.GroupClass},
			constant.NameAttribute:        {name},
			constant.GroupNumberAttribute: {strconv.Itoa(number)},
		},
	)

	if f != nil {
		return nil, f
	}

	return s.Group(name)
}
