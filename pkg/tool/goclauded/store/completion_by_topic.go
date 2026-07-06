package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"

func (s *Store) CompletionByTopic(
	sessionIdentifier string,
	topic string,
) (*completion.Completion, error) {
	var result completion.Completion
	r := s.database.Where(
		"session_identifier = ? AND topic = ?",
		sessionIdentifier,
		topic,
	).Limit(1).Find(&result)

	if r.Error != nil {
		return nil, r.Error
	}

	if r.RowsAffected == 0 {
		return nil, nil
	}

	return &result, nil
}
