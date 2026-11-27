package argument

type ReplyToThread struct {
	Team    string   `json:"team"`
	Channel string   `json:"channel"`
	Message string   `json:"message"`
	Text    string   `json:"text"`
	Mention []string `json:"mention"`
}
