package store

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
)

func (s *Store) EditEvent(
	identifier uint,
	value string,
) (*event.Event, error) {
	var existing event.Event

	if e := s.database.Where(
		"identifier = ?",
		identifier,
	).First(&existing).Error; e != nil {
		if relational.NotFound(e) {
			return nil, not_found.New("event", identifier)
		}

		return nil, e
	}

	key := editableKey(existing.Kind)

	if e := s.UpdateEventMetadata(identifier, key, value); e != nil {
		return nil, e
	}

	existing.Metadata = s.EventMetadataByEvent(identifier)

	return &existing, nil
}
