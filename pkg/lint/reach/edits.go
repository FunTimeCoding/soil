package reach

type Edits struct {
	Files      []string
	Patterns   []string
	Resolved   int
	Globbed    int
	Unresolved int
	Fixers     int
	Scripts    int
}
