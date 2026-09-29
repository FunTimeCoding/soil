package response

type NoteResponse struct {
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	Total       int    `json:"total"`
	PerPage     int    `json:"per_page"`
	Payload     []Note `json:"data"`
}
