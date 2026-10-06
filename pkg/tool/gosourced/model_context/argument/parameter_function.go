package argument

type ParameterFunction struct {
	PackagePath string   `json:"package_path"`
	Name        string   `json:"name"`
	Receiver    string   `json:"receiver"`
	Parameters  []string `json:"parameters"`
}
