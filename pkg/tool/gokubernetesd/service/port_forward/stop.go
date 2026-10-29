package port_forward

func (f *Forward) Stop() {
	f.once.Do(func() { close(f.stop) })
}
