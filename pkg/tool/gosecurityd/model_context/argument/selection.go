package argument

type Selection struct {
	Show   string            `json:"show"`
	Labels map[string]string `json:"labels"`
	Search string            `json:"search"`
}
