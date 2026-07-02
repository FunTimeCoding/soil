package target

import "fmt"

func FormatEntry(e *Entry) string {
	return fmt.Sprintf("%s: %s", e.Name, e.Value)
}
