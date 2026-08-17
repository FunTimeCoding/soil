package service

type profilePayload struct {
	Always   []memoryEntry `json:"always"`
	Relevant []memoryEntry `json:"relevant"`
}
