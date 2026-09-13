package service

import (
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/types/group"
)

func (s *Service) Groups() ([]*group.Group, error) {
	found, e := s.directory.Search(
		constant.GroupFilter,
		[]string{
			constant.NameAttribute,
			constant.GroupNumberAttribute,
			constant.MemberAttribute,
		},
	)

	if e != nil {
		return nil, e
	}

	var result []*group.Group

	for _, record := range found {
		result = append(
			result,
			group.New(
				first(record, constant.NameAttribute),
				groupNumber(record),
				record.Attributes[constant.MemberAttribute],
				record.DistinguishedName,
			),
		)
	}

	return result, nil
}
