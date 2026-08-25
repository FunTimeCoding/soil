package response

type Project struct {
	Identifier string `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
	Platform   string `json:"platform"`
}
