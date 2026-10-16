package plan

type Call struct {
	Expressions map[string]any `json:"expressions"`
	Module      *Module        `json:"module"`
}
