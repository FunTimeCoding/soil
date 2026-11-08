package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/label"

func (s *Store) RemoveLabel(
	sessionIdentifier string,
	key string,
	value string,
) ([]string, error) {
	old, e := s.GetLabel(sessionIdentifier, key)

	if e != nil {
		return nil, e
	}

	return old, s.database.Where(
		"session_identifier = ? AND key = ? AND value = ?",
		sessionIdentifier,
		key,
		value,
	).Delete(label.Stub()).Error
}
