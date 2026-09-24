package report

func (r *Report) Combined() float64 {
	var result float64

	for _, e := range r.Entries {
		result += e.Score
	}

	return result
}
