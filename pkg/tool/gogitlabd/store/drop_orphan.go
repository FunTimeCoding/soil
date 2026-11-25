package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Store) DropOrphan(k *watched_request.Key) error {
	remaining, e := s.ByRequest(k)

	if e != nil || len(remaining) > 0 {
		return e
	}

	return s.database.
		Where(
			"instance = ? AND project = ? AND request = ?",
			k.Instance,
			k.Project,
			k.Request,
		).
		Delete(snapshot.Stub()).Error
}
