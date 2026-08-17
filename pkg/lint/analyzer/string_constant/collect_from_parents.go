package string_constant

import "path/filepath"

func collectFromParents(
	result map[string][]knownConstant,
	directory string,
) {
	current := filepath.Dir(directory)

	for {
		collectFromConstantDirectory(result, current, "constant")

		if filepath.Base(current) == "pkg" {
			break
		}

		parent := filepath.Dir(current)

		if parent == current {
			break
		}

		current = parent
	}
}
