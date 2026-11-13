package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/types/checklist_item"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Service) WriteChecklist(
	key string,
	items []*checklist_item.Item,
) error {
	m, e := peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*field_map.Map, error) {
			return s.jira.FieldMap()
		},
	)

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

	return s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error { return s.jira.UpdateNative(raw) },
	)
}
