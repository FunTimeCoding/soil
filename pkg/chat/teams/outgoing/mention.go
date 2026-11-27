package outgoing

type Mention struct {
	Identifier int       `json:"id"`
	Text       string    `json:"mentionText"`
	Mentioned  Mentioned `json:"mentioned"`
}
