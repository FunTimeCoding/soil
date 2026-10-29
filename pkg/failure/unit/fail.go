package unit

func (f *Fixture) Fail(
	method any,
	e error,
) {
	f.failure.Set(method, e)
}
