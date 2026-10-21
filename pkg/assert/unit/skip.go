package unit

func (r *skipRecorder) Skip(...any) {
	r.skipped = true
}
