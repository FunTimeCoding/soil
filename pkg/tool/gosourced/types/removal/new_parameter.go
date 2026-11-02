package removal

func NewParameter(
	packagePath string,
	name string,
	receiver string,
	parameters []string,
) *Parameter {
	return &Parameter{
		PackagePath: packagePath,
		Name:        name,
		Receiver:    receiver,
		Parameters:  parameters,
	}
}
