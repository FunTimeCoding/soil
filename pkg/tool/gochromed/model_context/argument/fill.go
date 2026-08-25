package argument

type Fill struct {
	UID           string `json:"uid"`
	Value         string `json:"value"`
	TabIdentifier string `json:"tab_id"`
	Title         string `json:"title"`
	Locator       string `json:"url"`
	Snapshot      *bool  `json:"snapshot"`
	Direct        *bool  `json:"direct"`
}
