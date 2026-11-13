package pointer_dependency

func New() *Set {
	return &Set{
		paths:    make(map[string]bool),
		prefixes: make(map[string]bool),
	}
}
