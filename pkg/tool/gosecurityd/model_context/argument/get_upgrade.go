package argument

type GetUpgrade struct {
	Release string `json:"release"`
	Source  string `json:"source"`
}
