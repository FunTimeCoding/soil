package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Store) Delete(
	callsign string,
	k *watched_request.Key,
) (int, error) {
	r := s.database.
		Where(
			"callsign = ? AND instance = ? AND project = ? AND request = ?",
			callsign,
			k.Instance,
			k.Project,
			k.Request,
		).
		Delete(subscription.Stub())

	return int(r.RowsAffected), r.Error
}
