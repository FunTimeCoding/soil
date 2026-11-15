package failure_tester

func (f *Tester) Bravo() error {
	return f.failure.Caller()
}
