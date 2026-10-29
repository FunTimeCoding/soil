package unit

func (f *Fixture) Bravo() error {
	return f.failure.Caller()
}
