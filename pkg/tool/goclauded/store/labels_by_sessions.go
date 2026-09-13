package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
)

func (s *Store) LabelsBySessions(
	sessionIdentifiers []string,
) (map[string][]label.Label, error) {
	result := map[string][]label.Label{}

	if len(sessionIdentifiers) == 0 {
		return result, nil
	}

	var rows []label.Label
	e := s.database.Where("session_identifier IN ?", sessionIdentifiers).Order(
		constant.Key,
	).Find(
		&rows,
	).Error

	if e != nil {
		return nil, e
	}

	for _, row := range rows {
		result[row.SessionIdentifier] = append(
			result[row.SessionIdentifier],
			row,
		)
	}

	return result, nil
}
