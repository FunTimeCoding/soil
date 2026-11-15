package result

func NewUsage(
	packagePath string,
	typeName string,
	fields []string,
	methods []*UsageMethod,
	unloaded []*UsageUnloaded,
) *Usage {
	return &Usage{
		Package:  packagePath,
		Type:     typeName,
		Fields:   fields,
		Methods:  methods,
		Unloaded: unloaded,
	}
}
