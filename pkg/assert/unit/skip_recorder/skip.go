package skip_recorder

func (r *Recorder) Skip(...any) {
	r.skipped = true
}
