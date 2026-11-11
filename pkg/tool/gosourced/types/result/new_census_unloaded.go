package result

func NewCensusUnloaded(
	packagePath string,
	cause string,
	lost string,
) *CensusUnloaded {
	return &CensusUnloaded{Package: packagePath, Cause: cause, Lost: lost}
}
