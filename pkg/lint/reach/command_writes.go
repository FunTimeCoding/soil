package reach

type CommandWrites struct {
	Paths      []string
	Patterns   []string
	Unresolved bool
	Fixer      bool
	Script     bool
}
