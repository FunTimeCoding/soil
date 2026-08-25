package argument

type Screenshot struct {
	TabIdentifier string `json:"tab_id"`
	Title         string `json:"title"`
	Locator       string `json:"url"`
}
