package result

func NewLiterals(
	name string,
	total int,
	inside int,
	expected int,
	groups []*Group,
) *Literals {
	return &Literals{
		Type:     name,
		Total:    total,
		Inside:   inside,
		Expected: expected,
		Groups:   groups,
	}
}
