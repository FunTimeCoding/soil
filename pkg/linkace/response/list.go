package response

type List struct {
	Identifier  int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Visibility  int    `json:"visibility"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
