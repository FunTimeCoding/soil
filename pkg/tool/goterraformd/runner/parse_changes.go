package runner

import "strings"

func ParseChanges(output string) []string {
	var result []string

	for _, line := range strings.Split(output, "\n") {
		address, rest, found := strings.Cut(strings.TrimSpace(line), ": ")

		if !found {
			continue
		}

		if strings.HasPrefix(rest, "Creation complete") ||
			strings.HasPrefix(rest, "Modifications complete") ||
			strings.HasPrefix(rest, "Destruction complete") {
			result = append(result, address)
		}
	}

	return result
}
