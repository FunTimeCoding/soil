package gohook

import "slices"

func NewEntries(
	before []string,
	after []string,
) []string {
	var result []string

	for _, a := range after {
		if !slices.Contains(before, a) {
			result = append(result, a)
		}
	}

	return result
}
