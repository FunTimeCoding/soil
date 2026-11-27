package user

type User struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Mail       string `json:"mail,omitempty"`
	Principal  string `json:"principal,omitempty"`
	Title      string `json:"title,omitempty"`
}
