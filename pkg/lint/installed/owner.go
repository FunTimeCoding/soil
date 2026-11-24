package installed

func owner(
	sources map[string]string,
	name string,
	candidates []string,
) string {
	seen := make(map[string]bool)
	var implementers []string

	for _, c := range candidates {
		if i := implementer(sources, c, name); !seen[i] {
			seen[i] = true
			implementers = append(implementers, i)
		}
	}

	if len(implementers) == 1 {
		return implementers[0]
	}

	for _, i := range implementers {
		required := requires(sources[i])
		downstream := true

		for _, other := range implementers {
			if other != i && !required[other] {
				downstream = false

				break
			}
		}

		if downstream {
			return i
		}
	}

	return ""
}
