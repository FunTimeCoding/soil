package argument

type SendDirectMessage struct {
	User    string   `json:"user"`
	Text    string   `json:"text"`
	Mention []string `json:"mention"`
}
