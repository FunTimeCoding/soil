package request

type Subscribe struct {
	Events []string `json:"events"`
}
