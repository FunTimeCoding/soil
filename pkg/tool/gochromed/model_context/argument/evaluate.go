package argument

type Evaluate struct {
	Expression    string `json:"expression"`
	TabIdentifier string `json:"tab_id"`
	Title         string `json:"title"`
	Locator       string `json:"url"`
}
