package failure_tester

func (f *Tester) Fail(
	method any,
	e error,
) {
	f.failure.Set(method, e)
}
