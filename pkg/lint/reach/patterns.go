package reach

import "sort"

func patterns(directories map[string]bool) []string {
	var result []string

	for relative := range directories {
		result = append(result, directoryPattern(relative))
	}

	sort.Strings(result)

	return result
}
