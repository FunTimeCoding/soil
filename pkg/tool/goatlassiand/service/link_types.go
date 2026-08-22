package service

import (
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/link_type"
)

func (s *Service) LinkTypes() ([]link_type.Type, error) {
	_, body, e := s.jira.Basic().GetPath("rest/api/2/issueLinkType")

	if e != nil {
		return nil, e
	}

	var parsed linkTypeResponse
	notation.MustDecode(body, &parsed, true)

	return parsed.IssueLinkTypes, nil
}
