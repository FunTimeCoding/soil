package argument

type Create struct {
	Group  string   `json:"group"`
	Title  string   `json:"title"`
	Fields []string `json:"fields"`
}
