package census

func newReading(
	name string,
	file string,
) *reading {
	return &reading{
		name:      name,
		file:      file,
		calls:     map[string]bool{},
		functions: map[string]bool{},
		fields:    map[string]bool{},
		uses:      map[string]map[string]bool{},
		keys:      map[string]bool{},
		packages:  map[string]bool{},
		callers:   map[callerKey]int{},
	}
}
