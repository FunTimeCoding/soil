package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
	"gorm.io/gorm"
)

func (s *Store) SetLabel(
	sessionIdentifier string,
	key string,
	values []string,
) ([]string, error) {
	old, e := s.GetLabel(sessionIdentifier, key)

	if e != nil {
		return nil, e
	}

	return old, s.database.Transaction(
		func(t *gorm.DB) error {
			if f := t.Where(
				"session_identifier = ? AND key = ?",
				sessionIdentifier,
				key,
			).Delete(label.Stub()).Error; f != nil {
				return f
			}

			for _, v := range values {
				if g := t.Create(
					label.New(sessionIdentifier, key, v),
				).Error; g != nil {
					return g
				}
			}

			return nil
		},
	)
}
