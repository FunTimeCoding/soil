package response

type Reply struct {
	Identifier string   `json:"id"`
	Username   string   `json:"username"`
	Message    string   `json:"message"`
	CreateAt   string   `json:"create_at"`
	FileIds    []string `json:"file_ids,omitempty"`
}
