package mock_sink

import "github.com/funtimecoding/soil/pkg/generative/types/sink_event"

func (s *Sink) Push(
	content string,
	meta map[string]string,
) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.event = append(s.event, sink_event.Event{Content: content, Meta: meta})
}
