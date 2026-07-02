package target

import "fmt"

func SummarizeEntry(e *Entry) string {
	return fmt.Sprintf("[%s]", e.Name)
}
