package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event_metadata"
)

func (s *Store) UpsertEvent(
	sessionIdentifier string,
	kind string,
	actor string,
	metadata map[string]string,
) error {
	scope := metadata[constant.Topic]
	var existing event.Event
	query := s.database.Where(
		"session_identifier = ? AND kind = ?",
		sessionIdentifier,
		kind,
	)

	if scope != "" {
		query = query.Where(
			"identifier IN (SELECT event_identifier FROM event_metadata WHERE key = ? AND value = ?)",
			constant.Topic,
			scope,
		)
	}

	result := query.Limit(1).Find(&existing)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected > 0 {
		if e := s.database.Model(&existing).Update(
			"actor",
			actor,
		).Error; e != nil {
			return e
		}

		if e := s.database.Where(
			"event_identifier = ?",
			existing.Identifier,
		).Delete(event_metadata.Stub()).Error; e != nil {
			return e
		}

		return s.SetEventMetadata(existing.Identifier, metadata)
	}

	record := event.New(sessionIdentifier, kind, actor)

	if e := s.database.Create(record).Error; e != nil {
		return e
	}

	return s.SetEventMetadata(record.Identifier, metadata)
}
