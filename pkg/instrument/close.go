package instrument

func (i *Instrument) Close() {
	i.recorder.Flush()
	i.recorder.Close()
	i.reporter.Close()
}
