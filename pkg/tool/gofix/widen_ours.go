package gofix

import "slices"

func widenOurs(
	own func(string) bool,
	ours []string,
) func(string) bool {
	if len(ours) == 0 {
		return own
	}

	return func(modulePath string) bool {
		return own != nil && own(modulePath) || slices.Contains(ours, modulePath)
	}
}
