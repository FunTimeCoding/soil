package request_event

import "github.com/funtimecoding/soil/pkg/watch/digest/event"

func Visible(
	v []*Event,
	callsign string,
) []*event.Event {
	var result []*event.Event

	for _, e := range v {
		if e.Writer != "" && e.Writer == callsign {
			continue
		}

		result = append(result, e.Event)
	}

	return result
}
