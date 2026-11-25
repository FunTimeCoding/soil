package digest

import (
	"fmt"
	"strings"
)

func plural(
	count int,
	noun string,
) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, noun)
	}

	for _, suffix := range []string{"s", "x", "sh", "ch"} {
		if strings.HasSuffix(noun, suffix) {
			return fmt.Sprintf("%d %ses", count, noun)
		}
	}

	return fmt.Sprintf("%d %ss", count, noun)
}
