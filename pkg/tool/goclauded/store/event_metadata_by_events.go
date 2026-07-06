package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event_metadata"

func (s *Store) EventMetadataByEvents(
	identifiers []uint,
) map[uint]map[string]string {
	if len(identifiers) == 0 {
		return nil
	}

	var rows []event_metadata.EventMetadata
	s.database.Where(
		"event_identifier IN ?",
		identifiers,
	).Find(&rows)
	result := map[uint]map[string]string{}

	for _, row := range rows {
		if result[row.EventIdentifier] == nil {
			result[row.EventIdentifier] = map[string]string{}
		}

		result[row.EventIdentifier][row.Key] = row.Value
	}

	return result
}
