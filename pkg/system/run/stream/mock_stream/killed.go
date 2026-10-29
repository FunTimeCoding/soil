package mock_stream

func (s *Stream) Killed() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.killed
}
