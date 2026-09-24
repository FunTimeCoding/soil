package result

func (r *Result) AddUnplaced(path string) {
	r.Unplaced = append(r.Unplaced, path)
}
