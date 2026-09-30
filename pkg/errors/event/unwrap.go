package event

func (e *EventError) Unwrap() error {
	return e.Wrapped
}
