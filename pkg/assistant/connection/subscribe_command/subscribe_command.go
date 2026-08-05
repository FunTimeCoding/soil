package subscribe_command

type SubscribeCommand struct {
	Identifier uint64 `json:"id"`
	Type       string `json:"type"`
	Event      string `json:"event_type,omitempty"`
}
