package reading

func NewReading(
	name string,
	file string,
) *Reading {
	return &Reading{
		Name:      name,
		File:      file,
		Calls:     map[string]bool{},
		Functions: map[string]bool{},
		Fields:    map[string]bool{},
		Uses:      map[string]map[string]bool{},
		Keys:      map[string]bool{},
		Packages:  map[string]bool{},
		Callers:   map[CallerKey]int{},
	}
}
