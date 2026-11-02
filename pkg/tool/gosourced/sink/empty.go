package sink

func (s *Sink) Empty() bool {
	return len(s.rewrites) == 0 && len(s.operations) == 0
}
