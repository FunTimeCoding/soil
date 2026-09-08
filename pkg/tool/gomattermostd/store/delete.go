package store

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Store) Delete(callsign string, root string) (int, error) {
	r := s.database.
		Where("callsign = ? AND root_identifier = ?", callsign, root).
		Delete(subscription.Stub())

	return int(r.RowsAffected), r.Error
}
