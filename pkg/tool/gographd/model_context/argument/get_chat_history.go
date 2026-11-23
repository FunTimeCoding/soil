package argument

type GetChatHistory struct {
	Chat  string `json:"chat"`
	Limit int    `json:"limit"`
}
