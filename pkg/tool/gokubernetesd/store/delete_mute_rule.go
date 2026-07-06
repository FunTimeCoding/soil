package store

import "github.com/funtimecoding/soil/pkg/tool/gokubernetesd/store/mute_rule"

func (s *Store) DeleteMuteRule(identifier uint) error {
	return s.database.Delete(mute_rule.Stub(), identifier).Error
}
