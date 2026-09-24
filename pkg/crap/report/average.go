package report

func (r *Report) Average() float64 {
	if len(r.Entries) == 0 {
		return 0
	}

	return r.Combined() / float64(len(r.Entries))
}
