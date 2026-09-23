package mock_sink

func (s *Sink) Events() []Event {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return append([]Event{}, s.event...)
}
