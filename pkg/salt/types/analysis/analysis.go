package analysis

type Analysis struct {
	Reached          []string
	Possible         []string
	Unreached        []string
	Missing          []string
	UnreachedFiles   []string
	DeadEnvironments []string
}
