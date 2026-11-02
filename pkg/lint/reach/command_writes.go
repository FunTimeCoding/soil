package reach

type CommandWrites struct {
	Paths      []string
	Unresolved bool
	Fixer      bool
	Script     bool
}
