package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) applyFields(
	raw *jira.Issue,
	fields map[string]any,
) ([]string, error) {
	if len(fields) == 0 {
		return nil, nil
	}

	fieldMap, e := peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*field_map.Map, error) {
			return s.jira.FieldMap()
		},
	)

	if e != nil {
		return nil, e
	}

	var names []string

	for name, value := range fields {
		field := fieldMap.ByName(name)

		if field == nil {
			return nil, not_found.Format("unknown field: %s", name)
		}

		raw.Fields.Unknowns.Set(field.Key, value)
		names = append(names, name)
	}

	return names, nil
}
