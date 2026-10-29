package service

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"

func (s *Service) Entry(identifier uint) (*entry.Entry, error) {
	return s.store.Get(identifier)
}
