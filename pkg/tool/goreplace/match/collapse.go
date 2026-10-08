package match

import "strings"

func collapse(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
