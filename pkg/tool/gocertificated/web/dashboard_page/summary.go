package dashboard_page

import "fmt"

func Summary(
	authorities int,
	pending int,
	expiring int,
) []string {
	return []string{
		fmt.Sprintf("%d authorities", authorities),
		fmt.Sprintf("%d pending publication", pending),
		fmt.Sprintf("%d expiring within a month", expiring),
	}
}
