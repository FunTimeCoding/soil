package download

type Download struct {
	Identifier int    `json:"download_id"`
	Locator    string `json:"url"`
	Path       string `json:"filename"`
	State      string `json:"state"`
	Received   int64  `json:"received"`
	Total      int64  `json:"total"`
	Fail       string `json:"error"`
}
