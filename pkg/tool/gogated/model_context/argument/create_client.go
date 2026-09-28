package argument

type CreateClient struct {
	RedirectLocator string `json:"redirect_locator"`
	Scope           string `json:"scope"`
}
