package page_version

type Version struct {
	Number  int    `json:"number"`
	Message string `json:"message,omitempty"`
}
