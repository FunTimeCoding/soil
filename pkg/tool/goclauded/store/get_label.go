package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
)

func (s *Store) GetLabel(
	sessionIdentifier string,
	key string,
) ([]string, error) {
	var result []string

	return result, s.database.Model(label.Stub()).Where(
		"session_identifier = ? AND key = ?",
		sessionIdentifier,
		key,
	).Order(constant.Value).Pluck(constant.Value, &result).Error
}
