package supervisor

func (s *Supervisor) Done() <-chan struct{} {
	return s.allDone
}
