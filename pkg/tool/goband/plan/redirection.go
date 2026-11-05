package plan

type Redirection struct {
	Instance string   `json:"instance"`
	Pending  []string `json:"pending"`
}
