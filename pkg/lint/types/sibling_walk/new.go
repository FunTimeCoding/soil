package sibling_walk

func New() *Walk {
	return &Walk{Directories: make(map[string]bool)}
}
