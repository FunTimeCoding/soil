package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"

func (s *Store) EventsAfter(
	identifier uint,
	kinds []string,
	limit int,
) ([]event.Event, error) {
	g := s.database.Model(event.Stub()).Where("identifier > ?", identifier)

	if len(kinds) == 1 {
		g = g.Where("kind = ?", kinds[0])
	} else if len(kinds) > 1 {
		g = g.Where("kind IN ?", kinds)
	}

	var result []event.Event

	if e := g.Order("identifier ASC").Limit(limit).Find(&result).Error; e != nil {
		return nil, e
	}

	if len(result) == 0 {
		return result, nil
	}

	var identifiers []uint

	for _, v := range result {
		identifiers = append(identifiers, v.Identifier)
	}

	metadata := s.EventMetadataByEvents(identifiers)

	for i := range result {
		result[i].Metadata = metadata[result[i].Identifier]
	}

	return result, nil
}
