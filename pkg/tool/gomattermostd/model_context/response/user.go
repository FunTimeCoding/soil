package response

type User struct {
	Identifier string `json:"id"`
	Username   string `json:"username"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Nickname   string `json:"nickname"`
	Email      string `json:"email"`
	IsBot      bool   `json:"is_bot"`
}
