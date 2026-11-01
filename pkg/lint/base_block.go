package lint

import "strings"

func baseBlock(
	lines []string,
	baseLine int,
	number int,
) bool {
	if baseLine == 0 || number < baseLine {
		return false
	}

	for n := baseLine + 1; n <= number; n++ {
		line := lines[n-1]

		if !strings.HasPrefix(line, " ") && !strings.HasPrefix(line, "-") {
			return false
		}
	}

	return true
}
