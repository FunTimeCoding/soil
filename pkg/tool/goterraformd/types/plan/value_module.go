package plan

type ValueModule struct {
	Resources    []*ValueResource `json:"resources"`
	ChildModules []*ValueModule   `json:"child_modules"`
}
