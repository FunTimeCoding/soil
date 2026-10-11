package status

import "strings"

func FilterLog(s string) string {
	return strings.Split(s, "\n\n")[0]
}
