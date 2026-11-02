package reach

type Edits struct {
	Files      []string
	Resolved   int
	Unresolved int
	Fixers     int
	Scripts    int
}
