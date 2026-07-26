package target

import (
	"fmt"
	"strings"
)

// FormatName renders a labeled name.
func FormatName(name string) string {
	return fmt.Sprintf("name: %s", name)
}

func TrimName(name string) string {
	return strings.TrimSpace(name)
}

func PlainName(name string) string {
	return name
}
