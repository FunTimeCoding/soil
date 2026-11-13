package test_run

func (r *Run) named(action string) []string {
	var result []string

	for _, e := range r.events {
		if e.Action == action && e.Test != "" {
			result = append(result, e.Test)
		}
	}

	return sorted(result)
}
