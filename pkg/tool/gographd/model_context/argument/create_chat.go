package argument

type CreateChat struct {
	Members []string `json:"members"`
	Topic   string   `json:"topic"`
}
