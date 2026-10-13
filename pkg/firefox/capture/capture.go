package capture

type Capture struct {
	Identifier int    `json:"tab_id"`
	Locator    string `json:"url"`
	Title      string `json:"title"`
	Content    string `json:"content"`
}
