package pointer

import "strings"

func isBareLink(target string) bool {
	return target != "" &&
		!strings.ContainsAny(target, " \t") &&
		!strings.Contains(target, "://") &&
		!strings.HasPrefix(target, "#") &&
		!strings.ContainsAny(target, "<>*$") &&
		strings.ContainsRune(target, '.')
}
