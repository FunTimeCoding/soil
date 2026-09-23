package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func (s *Store) GroupedLabels(
	objectType string,
	identifiers []int32,
) (map[int32][]*label.Label, error) {
	result := map[int32][]*label.Label{}

	if len(identifiers) == 0 {
		return result, nil
	}

	var v []*label.Label
	e := s.database.
		Where(
			"object_type = ? AND object_identifier IN ?",
			objectType,
			identifiers,
		).
		Order(constant.KeyColumn).
		Find(&v).Error

	if e != nil {
		return nil, e
	}

	for _, l := range v {
		result[l.ObjectIdentifier] = append(result[l.ObjectIdentifier], l)
	}

	return result, nil
}
