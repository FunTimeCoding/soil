package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
)

func (s *Store) Snapshot(
	k *watched_request.Key,
) (*snapshot.Snapshot, bool, error) {
	var result []*snapshot.Snapshot
	e := s.database.
		Where(
			"instance = ? AND project = ? AND request = ?",
			k.Instance,
			k.Project,
			k.Request,
		).
		Limit(1).
		Find(&result).Error

	if e != nil || len(result) == 0 {
		return nil, false, e
	}

	return result[0], true, nil
}
