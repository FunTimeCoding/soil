package result

func NewCensusCaller(
	kind string,
	packagePath string,
	function string,
	count int,
) *CensusCaller {
	return &CensusCaller{
		Kind:     kind,
		Package:  packagePath,
		Function: function,
		Count:    count,
	}
}
