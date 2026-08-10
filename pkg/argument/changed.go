package argument

func (i *Instance) Changed(name string) bool {
	return i.flags.Changed(name)
}
