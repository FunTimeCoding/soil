package service

import "slices"

func searchPaths(path map[string]string) []string {
	var result []string

	for _, p := range path {
		for _, candidate := range []string{p, aliasPath(p)} {
			if candidate != "" && !slices.Contains(result, candidate) {
				result = append(result, candidate)
			}
		}
	}

	slices.Sort(result)

	return result
}
