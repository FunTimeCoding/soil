package store

import "github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"

func (s *Store) SaveSnapshot(v *snapshot.Snapshot) error {
	return s.database.Save(v).Error
}
