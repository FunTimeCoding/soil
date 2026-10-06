package fact

type Interface struct {
	Package string            `json:"package"`
	Name    string            `json:"name"`
	Methods map[string]string `json:"methods"`
}
