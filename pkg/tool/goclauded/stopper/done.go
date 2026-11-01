package stopper

func (s *Stopper) Done() <-chan struct{} {
	return s.done
}
