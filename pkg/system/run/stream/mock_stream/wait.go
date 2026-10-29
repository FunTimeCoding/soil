package mock_stream

func (s *Stream) Wait() error {
	return s.ending
}
