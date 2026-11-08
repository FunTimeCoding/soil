package result

func NewCensusMethod(
	name string,
	file string,
) *CensusMethod {
	return &CensusMethod{
		Name:   name,
		File:   file,
		Tables: NewCensusTables(nil, nil),
		Reach:  NewCensusTables(nil, nil),
	}
}
