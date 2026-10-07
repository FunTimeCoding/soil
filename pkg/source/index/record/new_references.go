package record

func NewReferences() *References {
	return &References{Targets: make(map[string][]*Site)}
}
