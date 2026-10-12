package imports

func Own(
	packagePath string,
	importPaths []string,
) string {
	best := -1
	result := ""
	tie := false

	for _, p := range importPaths {
		n := sharedSegments(packagePath, p)

		if n > best {
			best = n
			result = p
			tie = false

			continue
		}

		if n == best {
			tie = true
		}
	}

	if tie {
		return ""
	}

	return result
}
