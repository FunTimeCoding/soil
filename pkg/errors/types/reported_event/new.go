package reported_event

func New(error error) *Event {
	return &Event{Error: error}
}
