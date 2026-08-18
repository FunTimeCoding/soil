package output

func (r *Results) MarkPlanned() {
	for _, c := range r.Entries {
		if !c.Fixed {
			continue
		}

		c.Fixed = false
		c.Planned = true
	}
}
