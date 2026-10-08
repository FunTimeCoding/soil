package tab_response

type Response struct {
	Identifier string `json:"id"`
	Title      string `json:"title"`
	Locator    string `json:"url"`
	Type       string `json:"type,omitempty"`
	Parent     string `json:"parent,omitempty"`
}
