package response

type HistoryPost struct {
	Identifier     string   `json:"id"`
	Username       string   `json:"username"`
	Message        string   `json:"message"`
	CreateAt       string   `json:"create_at"`
	RootIdentifier *string  `json:"root_id,omitempty"`
	FileIds        []string `json:"file_ids,omitempty"`
}
