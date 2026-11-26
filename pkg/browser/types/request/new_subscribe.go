package request

func NewSubscribe(events []string) *Subscribe {
	return &Subscribe{Events: events}
}
