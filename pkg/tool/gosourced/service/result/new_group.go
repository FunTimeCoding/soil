package result

func NewGroup(
	shape string,
	exemplar string,
	locations []*Location,
) *Group {
	return &Group{Shape: shape, Exemplar: exemplar, Locations: locations}
}
