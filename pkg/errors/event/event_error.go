package event

type EventError struct {
	Type    string
	Raw     string
	Wrapped error
}
