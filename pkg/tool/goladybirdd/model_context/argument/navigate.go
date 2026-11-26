package argument

type Navigate struct {
	Locator       string `json:"url"`
	TabIdentifier string `json:"tab_id"`
}
