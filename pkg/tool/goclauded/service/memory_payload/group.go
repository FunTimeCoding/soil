package memory_payload

type Group struct {
	Parent   *Entry  `json:"parent"`
	Children []Entry `json:"children"`
}
