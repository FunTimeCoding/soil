package response

type Tag struct {
	Identifier int    `json:"id"`
	Name       string `json:"name"`
	Visibility int    `json:"visibility"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
