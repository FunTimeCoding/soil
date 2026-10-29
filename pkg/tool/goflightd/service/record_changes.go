package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/store/snapshot"
)

func (s *Service) recordChanges(
	kind string,
	values map[string]string,
) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	now := s.clock()

	for key, value := range values {
		seen := join.Empty(kind, constant.KeySeparator, key)

		if s.last[seen] == value {
			continue
		}

		if e := s.store.CreateSnapshot(
			snapshot.Snapshot{
				Time:  now,
				Kind:  kind,
				Key:   key,
				Value: value,
			},
		); e != nil {
			return e
		}
		s.last[seen] = value
	}

	return nil
}
