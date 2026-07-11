package service

func importPresent(
	names map[string]string,
	importPath string,
) bool {
	for _, p := range names {
		if p == importPath {
			return true
		}
	}

	return false
}
