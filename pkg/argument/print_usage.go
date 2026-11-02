package argument

func (i *Instance) PrintUsage() {
	if i.flags.Usage == nil {
		i.flags.PrintDefaults()

		return
	}

	i.flags.Usage()
}
