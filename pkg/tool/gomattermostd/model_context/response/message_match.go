package response

type MessageMatch struct {
	Identifier string   `json:"identifier"`
	Channel    string   `json:"channel"`
	Username   string   `json:"username"`
	Message    string   `json:"message"`
	CreateAt   string   `json:"create_at"`
	FileIds    []string `json:"file_ids,omitempty"`
}
