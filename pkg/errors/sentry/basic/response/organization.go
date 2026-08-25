package response

type Organization struct {
	Identifier string `json:"id"`
	Slug       string `json:"slug"`
	Name       string `json:"name"`
}
