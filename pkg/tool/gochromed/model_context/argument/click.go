package argument

type Click struct {
	UID           string `json:"uid"`
	TabIdentifier string `json:"tab_id"`
	Title         string `json:"title"`
	Locator       string `json:"url"`
	Snapshot      *bool  `json:"snapshot"`
}
