package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"

func (s *Store) CreateCompletion(
	sessionIdentifier string,
	name string,
	kind string,
	topic string,
	summary string,
) error {
	return s.database.Create(
		completion.New(
			sessionIdentifier,
			name,
			kind,
			topic,
			summary,
		),
	).Error
}
