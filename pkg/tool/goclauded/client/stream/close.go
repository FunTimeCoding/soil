package stream

import "github.com/funtimecoding/soil/pkg/errors"

func (s *Stream) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.body == nil {
		return
	}

	errors.LogClose(s.body)
	s.body = nil
}
