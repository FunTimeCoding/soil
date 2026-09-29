package response

type Link struct {
	Identifier    int            `json:"id"`
	Link          string         `json:"url"`
	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Icon          string         `json:"icon"`
	Status        int            `json:"status"`
	CheckDisabled bool           `json:"check_disabled"`
	Visibility    int            `json:"visibility"`
	CreatedAt     string         `json:"created_at"`
	UpdatedAt     string         `json:"updated_at"`
	Lists         []LinkRelation `json:"lists"`
	Tags          []LinkRelation `json:"tags"`
}
