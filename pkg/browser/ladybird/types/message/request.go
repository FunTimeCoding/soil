package message

type Request struct {
	Identifier int    `json:"id"`
	Method     string `json:"method"`
	Parameters any    `json:"params"`
}
