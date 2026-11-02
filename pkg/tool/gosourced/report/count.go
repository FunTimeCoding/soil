package report

func (r *Report) Count(class string) int {
	result := 0

	for _, e := range r.Entries {
		if e.Class == class {
			result++
		}
	}

	return result
}
