package ping_command

type PingCommand struct {
	Identifier uint64 `json:"id"`
	Type       string `json:"type"`
}
