package response

type Note struct {
	Identifier     int    `json:"id"`
	LinkIdentifier int    `json:"link_id"`
	Text           string `json:"note"`
	Visibility     int    `json:"visibility"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
}
