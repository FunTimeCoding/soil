package runner

func (r *Runner) seedLastSuccess() {
	v, e := r.store.LastSuccess()

	if e != nil || v.IsZero() {
		return
	}

	r.metrics.lastSuccess.Set(float64(v.Unix()))
}
