package response

type ChannelMatch struct {
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	LastPostAt  string `json:"last_post_at,omitempty"`
}
