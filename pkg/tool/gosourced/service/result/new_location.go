package result

func NewLocation(
	file string,
	line int,
	packagePath string,
) *Location {
	return &Location{
		File:    file,
		Line:    line,
		Package: packagePath,
	}
}
