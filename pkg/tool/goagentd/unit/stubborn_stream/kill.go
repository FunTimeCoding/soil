package stubborn_stream

func (s *Stream) Kill() error {
	if e := s.Stream.Kill(); e != nil {
		return e
	}

	return s.refusal
}
