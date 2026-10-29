package mock_system

import "github.com/funtimecoding/soil/pkg/face"

func (s *System) Stream(...string) face.Stream {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.stream
}
