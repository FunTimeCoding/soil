package port_forward

func (f *Forward) Stopping() <-chan struct{} {
	return f.stop
}
