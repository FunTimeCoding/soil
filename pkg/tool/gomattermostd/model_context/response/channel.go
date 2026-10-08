package response

type Channel struct {
	Identifier  string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name"`
	Type        string `json:"type"`
	Purpose     string `json:"purpose"`
	Header      string `json:"header"`
}
