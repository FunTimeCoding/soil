package memory_payload

type Profile struct {
	Always   []Entry `json:"always"`
	Relevant []Entry `json:"relevant"`
}
