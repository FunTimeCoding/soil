package result

func NewUsageUnloaded(
	packagePath string,
	cause string,
	lost string,
) *UsageUnloaded {
	return &UsageUnloaded{Package: packagePath, Cause: cause, Lost: lost}
}
