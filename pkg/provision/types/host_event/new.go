package host_event

func New(
	host string,
	action string,
) *Event {
	return &Event{Host: host, Action: action}
}
