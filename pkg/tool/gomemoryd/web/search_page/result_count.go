package search_page

import "fmt"

func resultCount(count int) string {
	if count == 1 {
		return "1 result"
	}

	return fmt.Sprintf("%d results", count)
}
