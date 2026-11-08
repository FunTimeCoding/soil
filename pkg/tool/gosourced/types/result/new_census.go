package result

func NewCensus(
	packagePath string,
	typeName string,
	fields []string,
	methods []*CensusMethod,
) *Census {
	return &Census{
		Package: packagePath,
		Type:    typeName,
		Fields:  fields,
		Methods: methods,
	}
}
