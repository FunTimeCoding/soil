package allowlist

func (r *Rule) Match(words []string) bool {
	if len(words) < len(r.Words) {
		return false
	}

	if !r.Prefix && len(words) != len(r.Words) {
		return false
	}

	for i, w := range r.Words {
		if words[i] != w {
			return false
		}
	}

	return true
}
