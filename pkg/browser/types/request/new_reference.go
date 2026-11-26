package request

func NewReference(shared string) *Reference {
	return &Reference{SharedIdentifier: shared}
}
