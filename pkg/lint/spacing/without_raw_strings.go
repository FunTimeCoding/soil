package spacing

import "strings"

func withoutRawStrings(line string) string {
	var result strings.Builder
	inside := false

	for _, r := range line {
		if r == '`' {
			inside = !inside

			continue
		}

		if !inside {
			result.WriteRune(r)
		}
	}

	return result.String()
}
