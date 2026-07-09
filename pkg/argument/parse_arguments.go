package argument

func (i *Instance) ParseArguments(arguments []string) error {
	return i.flags.Parse(arguments)
}
