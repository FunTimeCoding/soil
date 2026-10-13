package argument

type DownloadFile struct {
	Locator string `json:"url"`
	Path    string `json:"filename"`
}
