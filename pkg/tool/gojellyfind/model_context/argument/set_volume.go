package argument

type SetVolume struct {
	SessionIdentifier string `json:"session_id"`
	Level             int    `json:"level"`
}
