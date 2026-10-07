package mock_sink

import "github.com/funtimecoding/soil/pkg/generative/types/sink_event"

func (s *Sink) Events() []sink_event.Event {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return append([]sink_event.Event{}, s.event...)
}
