package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"
)

func (s *Store) LabelsBySession(sessionIdentifier string) ([]label.Label, error) {
	var result []label.Label

	return result, s.database.Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Order(constant.Key).Find(&result).Error
}
