package relay_tester

func (r *Relay) Address() string {
	return r.listener.Addr().String()
}
