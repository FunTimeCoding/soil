package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"
)

func (s *Store) GetEvent(identifier uint) *event.Event {
	var e event.Event
	result := s.database.Where("identifier = ?", identifier).Limit(1).Find(&e)
	errors.PanicOnError(result.Error)

	if result.RowsAffected == 0 {
		return nil
	}

	e.Metadata = s.EventMetadataByEvent(identifier)

	return &e
}
