package model_context

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
)

func (s *Server) readChecklist(key string) ([]*checklist_item.Item, error) {
	i, e := s.jira.Issue(key)

	if e != nil {
		return nil, e
	}

	value := i.CustomValue(constant.ChecklistField)

	if value == "" ||
		value == issue.NilValue ||
		value == issue.UnknownField ||
		value == issue.UnknownValue {
		return nil, nil
	}

	return ParseChecklist(value), nil
}
