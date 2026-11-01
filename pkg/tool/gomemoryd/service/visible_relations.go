package service

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/record"

func (s *Service) VisibleRelations() ([]record.RelationOverview, error) {
	relations, e := s.store.ListRelations()

	if e != nil {
		return nil, e
	}

	hidden, f := s.HiddenIdentifiers()

	if f != nil {
		return nil, f
	}

	result := make([]record.RelationOverview, 0, len(relations))

	for _, v := range relations {
		if hidden[v.SourceIdentifier] || hidden[v.TargetIdentifier] {
			continue
		}

		result = append(result, v)
	}

	return result, nil
}
