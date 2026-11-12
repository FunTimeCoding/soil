package qualifier

type Qualifier struct {
	Imports map[string]string
	Own     map[string]bool
	Path    string
}
