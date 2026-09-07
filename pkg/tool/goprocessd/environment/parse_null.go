package environment

import (
	"bufio"
	"strings"
)

func parseNull(output string) map[string]string {
	result := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(output))
	scanner.Split(SplitNull)

	for scanner.Scan() {
		key, value, found := strings.Cut(scanner.Text(), "=")

		if found {
			result[key] = value
		}
	}

	return result
}
