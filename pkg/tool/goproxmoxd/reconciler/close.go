package reconciler

func (r *Reconciler) Close() {
	r.cancel()
}
