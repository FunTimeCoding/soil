package stream

func (s *Stream) Kill() error {
	return s.command.Process.Kill()
}
