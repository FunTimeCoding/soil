package mock_sink

import (
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/generative/types/sink_event"
)

func (s *Sink) Kind(kind string) []sink_event.Event {
	var result []sink_event.Event

	for _, event := range s.Events() {
		if event.Meta[constant.ChannelKindMeta] == kind {
			result = append(result, event)
		}
	}

	return result
}
