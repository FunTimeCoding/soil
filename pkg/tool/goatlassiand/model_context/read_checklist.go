package model_context

import (
	atlassian "github.com/funtimecoding/soil/pkg/atlassian/constant"
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
		value == atlassian.JiraNilValue ||
		value == atlassian.JiraUnknownField ||
		value == atlassian.JiraUnknownValue {
		return nil, nil
	}

	return ParseChecklist(value), nil
}
