package argument

type GetDebianStatus struct {
	Identifier string `json:"identifier"`
	Release    string `json:"release"`
}
