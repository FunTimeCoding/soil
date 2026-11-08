package census

type reading struct {
	name       string
	file       string
	calls      map[string]bool
	functions  map[string]bool
	fields     map[string]bool
	uses       map[string]map[string]bool
	keys       map[string]bool
	packages   map[string]bool
	atomically bool
	step       bool
	callers    map[callerKey]int
}
