package mock_stream

func (s *Stream) Kill() error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.killed++

	return s.writer.Close()
}
