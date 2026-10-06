package instrument

func (i *Instrument) BeginCommand(name string) {
	i.command = name
}
