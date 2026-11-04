package argument

type AdditionFunction struct {
	PackagePath string `json:"package_path"`
	Name        string `json:"name"`
	Receiver    string `json:"receiver"`
	Parameter   string `json:"parameter"`
	Type        string `json:"type"`
	Position    int    `json:"position"`
	Argument    string `json:"argument"`
}
