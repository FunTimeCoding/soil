package mock_sink

import "github.com/funtimecoding/soil/pkg/generative/constant"

func (s *Sink) Kind(kind string) []Event {
	var result []Event

	for _, event := range s.Events() {
		if event.Meta[constant.ChannelKindMeta] == kind {
			result = append(result, event)
		}
	}

	return result
}
