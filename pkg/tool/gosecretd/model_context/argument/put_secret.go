package argument

type PutSecret struct {
	Path   string            `json:"path"`
	Values map[string]string `json:"values"`
}
