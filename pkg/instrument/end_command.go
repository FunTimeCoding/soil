package instrument

func (i *Instrument) EndCommand(outcome string) {
	if i.command != "" {
		i.record(i.command, outcome)
		i.command = ""
	}

	i.recorder.Flush()
}
