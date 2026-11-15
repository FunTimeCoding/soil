package failure_tester

func (f *Tester) Alfa() error {
	return f.failure.Caller()
}
