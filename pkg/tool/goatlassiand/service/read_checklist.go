package service

import (
	atlassianConstant "github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
)

func (s *Service) ReadChecklist(key string) ([]*checklist_item.Item, error) {
	i, e := s.Issue(key)

	if e != nil {
		return nil, e
	}

	value := i.CustomValue(constant.ChecklistField)

	if value == "" ||
		value == atlassianConstant.JiraNilValue ||
		value == atlassianConstant.JiraUnknownField ||
		value == atlassianConstant.JiraUnknownValue {
		return nil, nil
	}

	return checklist_item.Parse(value), nil
}
