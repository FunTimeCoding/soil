package digest

import "fmt"

func plural(
	count int,
	noun string,
) string {
	if count == 1 {
		return fmt.Sprintf("%d %s", count, noun)
	}

	return fmt.Sprintf("%d %ss", count, noun)
}
