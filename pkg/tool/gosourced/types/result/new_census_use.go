package result

func NewCensusUse(
	typeName string,
	methods []string,
) *CensusUse {
	return &CensusUse{Type: typeName, Methods: methods}
}
