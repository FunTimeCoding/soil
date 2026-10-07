package message

type Request struct {
	Method     string `json:"method"`
	Parameters any    `json:"params"`
	Identifier int    `json:"id"`
}
