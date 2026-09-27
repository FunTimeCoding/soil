package stream_event

type Event struct {
	Identifier        uint              `json:"identifier"`
	SessionIdentifier string            `json:"session_identifier"`
	Kind              string            `json:"kind"`
	Actor             string            `json:"actor"`
	Created           string            `json:"created"`
	Metadata          map[string]string `json:"metadata,omitempty"`
}
