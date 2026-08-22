package unreadable_body

type UnreadableBodyError struct {
	Message string
	Wrapped error
}
