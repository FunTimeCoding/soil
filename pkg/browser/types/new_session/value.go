package new_session

type Value struct {
	SessionIdentifier string        `json:"sessionId"`
	Capabilities      *Capabilities `json:"capabilities"`
}
