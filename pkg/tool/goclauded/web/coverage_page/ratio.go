package coverage_page

import "fmt"

func ratio(
	used int,
	registered int,
) string {
	if registered == 0 {
		return "-"
	}

	return fmt.Sprintf("%d/%d (%d%%)", used, registered, used*100/registered)
}
