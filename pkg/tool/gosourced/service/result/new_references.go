package result

func NewReferences(
	symbol string,
	locations []*Location,
) *References {
	return &References{
		Symbol:    symbol,
		Total:     len(locations),
		Locations: locations,
	}
}
