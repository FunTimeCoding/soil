package argument

func (i *Instance) Float(
	name string,
	value float64,
	usage string,
) {
	i.flags.Float64(name, value, usage)
}
