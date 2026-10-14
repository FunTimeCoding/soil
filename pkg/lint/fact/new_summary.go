package fact

func NewSummary() *Summary {
	return &Summary{
		Arranging: make(map[string][]Condition),
		Delegates: make(map[string][]string),
		Disposes:  make(map[string][]int),
		Forwards:  make(map[string][]Forward),
		Releases:  make(map[string][]Release),
		SpentVia:  make(map[string][]string),
		LendsVia:  make(map[string][]string),
		Wraps:     make(map[string][]Forward),
	}
}
