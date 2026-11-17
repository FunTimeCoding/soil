package tool_tester

type Schema struct {
	Required   []string             `json:"required"`
	Properties map[string]*Property `json:"properties"`
}
