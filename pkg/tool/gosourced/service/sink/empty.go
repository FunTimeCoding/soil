package sink

func (s *Sink) Empty() bool {
	return len(s.operations) == 0
}
