package addition

func NewParameter(
	packagePath string,
	name string,
	receiver string,
	parameter string,
	typeSource string,
	position int,
	argument string,
) *Parameter {
	return &Parameter{
		PackagePath: packagePath,
		Name:        name,
		Receiver:    receiver,
		Parameter:   parameter,
		Type:        typeSource,
		Position:    position,
		Argument:    argument,
	}
}
