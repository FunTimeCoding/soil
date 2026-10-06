package result

func NewConstructor(
	name string,
	constructor string,
	file string,
	parameters []string,
) *Constructor {
	return &Constructor{
		Type:       name,
		Name:       constructor,
		File:       file,
		Parameters: parameters,
	}
}
