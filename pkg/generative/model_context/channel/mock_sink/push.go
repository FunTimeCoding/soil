package mock_sink

func (s *Sink) Push(
	content string,
	meta map[string]string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.event = append(s.event, Event{Content: content, Meta: meta})
}
