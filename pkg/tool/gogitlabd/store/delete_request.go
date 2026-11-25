package store

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/snapshot"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/store/subscription"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/watched_request"
	"gorm.io/gorm"
)

func (s *Store) DeleteRequest(k *watched_request.Key) error {
	return s.database.Transaction(
		func(t *gorm.DB) error {
			match := "instance = ? AND project = ? AND request = ?"

			if e := t.Where(
				match,
				k.Instance,
				k.Project,
				k.Request,
			).Delete(subscription.Stub()).Error; e != nil {
				return e
			}

			return t.Where(
				match,
				k.Instance,
				k.Project,
				k.Request,
			).Delete(snapshot.Stub()).Error
		},
	)
}
