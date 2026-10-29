package unit

func (f *Fixture) Alfa() error {
	return f.failure.Caller()
}
