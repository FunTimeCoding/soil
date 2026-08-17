package service

type groupPayload struct {
	Parent   *memoryEntry  `json:"parent"`
	Children []memoryEntry `json:"children"`
}
