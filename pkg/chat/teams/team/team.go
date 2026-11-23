package team

type Team struct {
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}
