package mock_system

import "github.com/funtimecoding/soil/pkg/system/run/stream/mock_stream"

func (s *System) Serve(v *mock_stream.Stream) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.stream = v
}
