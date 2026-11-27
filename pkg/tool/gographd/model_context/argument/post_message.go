package argument

type PostMessage struct {
	Team    string   `json:"team"`
	Channel string   `json:"channel"`
	Text    string   `json:"text"`
	Mention []string `json:"mention"`
}
