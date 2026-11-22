package runner

import "strings"

func listed(
	output string,
	host string,
) bool {
	for _, l := range strings.Split(output, "\n") {
		if strings.TrimSpace(l) == host {
			return true
		}
	}

	return false
}
