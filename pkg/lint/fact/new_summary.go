package fact

func NewSummary() *Summary {
	return &Summary{Delegates: make(map[string][]string)}
}
