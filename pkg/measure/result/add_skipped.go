package result

func (r *Result) AddSkipped(path string) {
	r.Skipped = append(r.Skipped, path)
}
