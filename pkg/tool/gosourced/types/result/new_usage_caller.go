package result

func NewUsageCaller(
	kind string,
	packagePath string,
	function string,
	count int,
) *UsageCaller {
	return &UsageCaller{
		Kind:     kind,
		Package:  packagePath,
		Function: function,
		Count:    count,
	}
}
