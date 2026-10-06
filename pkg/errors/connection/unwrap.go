package connection

func (f *Failure) Unwrap() error {
	return f.class
}
