package reported_event

type Event struct {
	Error   error
	Context map[string]any
}
