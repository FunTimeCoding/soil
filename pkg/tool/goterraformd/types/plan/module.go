package plan

type Module struct {
	Resources   []*Resource        `json:"resources"`
	ModuleCalls map[string]*Call   `json:"module_calls"`
	Outputs     map[string]*Output `json:"outputs"`
}
