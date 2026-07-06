package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/event"

func (s *Store) LogEvent(
	sessionIdentifier string,
	kind string,
	actor string,
	metadata map[string]string,
) error {
	record := event.New(sessionIdentifier, kind, actor)

	if e := s.database.Create(record).Error; e != nil {
		return e
	}

	return s.SetEventMetadata(record.Identifier, metadata)
}
