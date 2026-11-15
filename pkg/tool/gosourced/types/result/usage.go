package result

type Usage struct {
	Package  string           `json:"package"`
	Type     string           `json:"type"`
	Fields   []string         `json:"fields"`
	Methods  []*UsageMethod   `json:"methods"`
	Unloaded []*UsageUnloaded `json:"unloaded"`
}
