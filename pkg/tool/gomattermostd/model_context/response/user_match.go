package response

type UserMatch struct {
	Identifier string `json:"identifier"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Nickname   string `json:"nickname,omitempty"`
	Email      string `json:"email,omitempty"`
}
