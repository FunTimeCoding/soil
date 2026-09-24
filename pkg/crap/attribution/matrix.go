package attribution

type Matrix struct {
	Tests  []string            `json:"tests"`
	Covers map[string][]string `json:"covers"`
}
