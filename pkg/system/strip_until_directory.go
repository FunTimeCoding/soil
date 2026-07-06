package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func StripUntilDirectory(
	s string,
	directory string,
) string {
	parts := strings.SplitN(s, directory, 2)

	if len(parts) < 2 {
		return s
	}

	return fmt.Sprintf("%s%s%s", separator.Slash, directory, parts[1])
}
