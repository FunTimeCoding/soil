package pointer

func (r *Resolver) anchored(
	bases []string,
	candidate string,
) (string, bool) {
	for _, base := range bases {
		segment, full := Anchor(base, candidate)

		if segment == "" {
			continue
		}

		if r.Exists(segment) || r.SiblingExists(segment) {
			return full, true
		}
	}

	return "", false
}
