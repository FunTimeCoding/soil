package service

import "github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/link_type"

func (s *Service) LinkTypes() ([]link_type.Type, error) {
	var parsed linkTypeResponse

	if e := s.jira.Basic().GetPath(
		"rest/api/2/issueLinkType",
		&parsed,
	); e != nil {
		return nil, e
	}

	return parsed.IssueLinkTypes, nil
}
