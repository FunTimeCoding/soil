package reporter

func (r *Reporter) Close() {
	if r.client != nil {
		r.client.Close()
	}

	r.pool.Close()
}
