package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Service) WriteChecklist(
	key string,
	items []*checklist_item.Item,
) error {
	m, e := s.jira.FieldMap()

	if e != nil {
		return e
	}

	field := m.ByName(constant.ChecklistField)

	if field == nil {
		return not_found.Format("checklist field not found")
	}

	raw := issue.Raw(key)
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)
	raw.Fields.Unknowns.Set(field.Key, checklist_item.Format(items))

	return s.jira.UpdateNative(raw)
}
