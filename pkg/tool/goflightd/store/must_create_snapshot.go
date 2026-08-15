package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
)

func (s *Store) MustCreateSnapshot(v snapshot.Snapshot) {
	errors.PanicOnError(s.CreateSnapshot(v))
}
