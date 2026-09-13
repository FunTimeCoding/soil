package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/store/session"

func (s *Store) MarkClosed(
	identifier string,
	reason string,
) error {
	closed := s.clock()

	return s.database.Model(session.Stub()).Where(
		"identifier = ?",
		identifier,
	).Updates(
		map[string]any{"closed_at": &closed, "closed_reason": reason},
	).Error
}
