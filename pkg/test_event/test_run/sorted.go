package test_run

import "slices"

func sorted(names []string) []string {
	slices.Sort(names)

	return names
}
