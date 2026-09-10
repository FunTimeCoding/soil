package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event_metadata"
)

func (s *Store) CountSessionEventMetadata(
	sessionIdentifier string,
) (int64, error) {
	var result int64

	return result, s.database.Model(event_metadata.Stub()).Where(
		"event_identifier IN (?)",
		s.database.Model(event.Stub()).Select(constant.Identifier).Where(
			"session_identifier = ?",
			sessionIdentifier,
		),
	).Count(&result).Error
}
