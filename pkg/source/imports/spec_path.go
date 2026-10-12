package imports

import "strings"

func specPath(line string) string {
	if i := strings.Index(line, `"`); i >= 0 {
		return line[i:]
	}

	return line
}
