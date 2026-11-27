package outgoing

type Chat struct {
	Kind    string   `json:"chatType"`
	Topic   string   `json:"topic,omitempty"`
	Members []Member `json:"members"`
}
