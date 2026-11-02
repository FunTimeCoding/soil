package registry

func (r *Registry) Close() {
	for _, f := range r.All() {
		f.Stop()
	}

	r.running.Wait()
}
