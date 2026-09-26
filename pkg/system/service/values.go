package service

import "slices"

func values(v map[string]string) []string {
	var result []string

	for _, s := range v {
		if s != "" && !slices.Contains(result, s) {
			result = append(result, s)
		}
	}

	slices.Sort(result)

	return result
}
