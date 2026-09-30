package event

func New(
	kind string,
	raw string,
	wrapped error,
) *EventError {
	return &EventError{Type: kind, Raw: raw, Wrapped: wrapped}
}
