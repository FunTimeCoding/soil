package imports

import "strings"

func ConstantKeeper(
	packagePath string,
	importPaths []string,
) string {
	result := ""

	for _, p := range importPaths {
		if !IsOwn(packagePath, p) {
			continue
		}

		if strings.Count(p, "/") > strings.Count(result, "/") {
			result = p
		}
	}

	return result
}
