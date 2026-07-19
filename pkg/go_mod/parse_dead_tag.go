package go_mod

func ParseDeadTag(stderr string) (string, string) {
	m := deadTagPattern.FindStringSubmatch(stderr)

	if m == nil {
		return "", ""
	}

	return m[1], m[2]
}
