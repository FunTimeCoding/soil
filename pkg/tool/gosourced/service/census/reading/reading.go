package reading

type Reading struct {
	Name       string
	File       string
	Calls      map[string]bool
	Functions  map[string]bool
	Fields     map[string]bool
	Uses       map[string]map[string]bool
	Keys       map[string]bool
	Packages   map[string]bool
	Atomically bool
	Step       bool
	Callers    map[CallerKey]int
}
