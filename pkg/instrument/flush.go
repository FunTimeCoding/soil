package instrument

func (i *Instrument) Flush(v any) {
	i.recorder.Flush()
	i.reporter.RecoverFlush(v)
}
