package strings

func filterMembership(
	a []string,
	b []string,
	member bool,
) []string {
	buffer := make(map[string]struct{}, len(b))

	for _, x := range b {
		buffer[x] = struct{}{}
	}

	result := make([]string, 0)

	for _, x := range a {
		if _, okay := buffer[x]; okay == member {
			result = append(result, x)
		}
	}

	return result
}
