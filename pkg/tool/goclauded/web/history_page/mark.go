package history_page

import "fmt"

func mark(identifier uint) string {
	return fmt.Sprintf("event-%d", identifier)
}
