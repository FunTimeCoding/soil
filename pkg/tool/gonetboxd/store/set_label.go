package store

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"

func (s *Store) SetLabel(
	objectType string,
	objectIdentifier int32,
	key string,
	value string,
) (*label.Label, error) {
	result := label.New()
	e := s.database.
		Where(
			"object_type = ? AND object_identifier = ? AND key = ?",
			objectType,
			objectIdentifier,
			key,
		).
		Assign(map[string]any{"value": value}).
		FirstOrCreate(
			result,
			label.Label{
				ObjectType:       objectType,
				ObjectIdentifier: objectIdentifier,
				Key:              key,
			},
		).Error

	return result, e
}
