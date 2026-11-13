package pointer_dependency

type Set struct {
	paths    map[string]bool
	prefixes map[string]bool
	Tests    bool
	Outside  bool
}
