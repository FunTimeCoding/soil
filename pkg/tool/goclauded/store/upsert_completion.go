package store

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/completion"
)

func (s *Store) UpsertCompletion(
	sessionIdentifier string,
	name string,
	kind string,
	topic string,
	message string,
) (*completion.Completion, error) {
	var existing completion.Completion
	result := s.database.Where(
		"session_identifier = ? AND topic = ?",
		sessionIdentifier,
		topic,
	).Limit(1).Find(&existing)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected > 0 {
		e := s.database.Model(&existing).Updates(
			map[string]any{
				"name":                 name,
				"kind":                 kind,
				constant.SummaryColumn: message,
			},
		).Error

		if e != nil {
			return nil, e
		}

		return &existing, nil
	}

	var count int64
	s.database.Model(completion.Stub()).Where(
		"session_identifier = ?",
		sessionIdentifier,
	).Count(&count)
	c := completion.New(sessionIdentifier, name, kind, topic, message)
	c.Sequence = int(count) + 1

	if e := s.database.Create(c).Error; e != nil {
		return nil, e
	}

	return c, nil
}
