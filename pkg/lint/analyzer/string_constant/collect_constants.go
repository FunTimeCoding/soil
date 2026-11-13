package string_constant

import "path/filepath"

func collectConstants(packageDirectory string) map[string][]KnownConstant {
	result := make(map[string][]KnownConstant)

	for _, path := range ConstantFiles(packageDirectory) {
		qualifier := "constant"

		if filepath.Dir(path) == packageDirectory {
			qualifier = ""
		}

		parseConstants(result, path, qualifier)
	}

	return result
}
