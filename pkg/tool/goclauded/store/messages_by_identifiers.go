package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/message"
)

func (s *Store) MessagesByIdentifiers(
	identifiers []uint,
) ([]*message.Message, error) {
	var result []*message.Message

	if len(identifiers) == 0 {
		return result, nil
	}

	return result, s.database.Where(
		"identifier IN ?",
		identifiers,
	).Order(constant.Identifier).Find(&result).Error
}
