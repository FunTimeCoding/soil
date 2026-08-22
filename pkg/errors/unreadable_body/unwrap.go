package unreadable_body

func (e *UnreadableBodyError) Unwrap() error {
	return e.Wrapped
}
