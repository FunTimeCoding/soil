package argument

type CreateDevice struct {
	Name   string   `json:"name"`
	Role   string   `json:"role"`
	Type   string   `json:"type"`
	Site   string   `json:"site"`
	Tenant string   `json:"tenant"`
	Tags   []string `json:"tags"`
}
