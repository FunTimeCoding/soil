package response

type Context struct {
	Identifier string     `json:"context"`
	Locator    string     `json:"url"`
	Children   []*Context `json:"children"`
}
