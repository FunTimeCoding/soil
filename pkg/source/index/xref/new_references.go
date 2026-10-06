package xref

func NewReferences() *References {
	return &References{Targets: make(map[string][]*Site)}
}
