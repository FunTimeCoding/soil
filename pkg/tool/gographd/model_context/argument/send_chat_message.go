package argument

type SendChatMessage struct {
	Chat    string   `json:"chat"`
	Text    string   `json:"text"`
	Mention []string `json:"mention"`
}
