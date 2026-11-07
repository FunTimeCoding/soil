package policy

func (w *Walk) allowlisted(words []string) bool {
	for _, r := range w.rules {
		if r.Match(words) {
			return true
		}
	}

	return false
}
