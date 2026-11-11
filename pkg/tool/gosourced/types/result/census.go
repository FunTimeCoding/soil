package result

type Census struct {
	Package  string            `json:"package"`
	Type     string            `json:"type"`
	Fields   []string          `json:"fields"`
	Methods  []*CensusMethod   `json:"methods"`
	Unloaded []*CensusUnloaded `json:"unloaded"`
}
