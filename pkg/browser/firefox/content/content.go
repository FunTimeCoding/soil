package content

type Content struct {
	Identifier int    `json:"tab_id"`
	Locator    string `json:"url"`
	Title      string `json:"title"`
	Text       string `json:"text"`
	Mode       string `json:"mode"`
}
