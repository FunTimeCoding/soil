package installed

func relevant(
	main string,
	directories map[string]string,
	dependencies map[string][]string,
) map[string]bool {
	result := make(map[string]bool)
	own, okay := directories[main]

	if !okay {
		return result
	}

	result[own] = true

	for _, d := range dependencies[main] {
		if local, found := directories[d]; found {
			result[local] = true
		}
	}

	return result
}
