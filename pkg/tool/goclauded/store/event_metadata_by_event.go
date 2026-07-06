package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event_metadata"

func (s *Store) EventMetadataByEvent(eventIdentifier uint) map[string]string {
	var rows []event_metadata.EventMetadata
	s.database.Where(
		"event_identifier = ?",
		eventIdentifier,
	).Find(&rows)

	if len(rows) == 0 {
		return nil
	}

	result := map[string]string{}

	for _, row := range rows {
		result[row.Key] = row.Value
	}

	return result
}
