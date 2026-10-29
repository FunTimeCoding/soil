package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/types/change"
)

func (s *Service) Update(
	identifier uint,
	c *change.Change,
) (*entry.Entry, error) {
	result, e := s.store.Get(identifier)

	if e != nil {
		return nil, e
	}

	c.Apply(result)

	if f := check(result); f != nil {
		return nil, f
	}

	if f := s.store.Update(result); f != nil {
		return nil, f
	}

	s.notifier.Notify()

	return result, nil
}
