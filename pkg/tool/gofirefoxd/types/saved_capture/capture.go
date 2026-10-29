package saved_capture

type Capture struct {
	Path  string `json:"path"`
	Title string `json:"tab"`
	Size  int    `json:"size"`
}
