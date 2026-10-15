package relay_tester

func (r *Relay) Freeze() {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, f := range r.frozen {
		f.Store(true)
	}
}
