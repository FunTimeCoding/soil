package response

type ListResponse struct {
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	Total       int    `json:"total"`
	PerPage     int    `json:"per_page"`
	Payload     []List `json:"data"`
}
