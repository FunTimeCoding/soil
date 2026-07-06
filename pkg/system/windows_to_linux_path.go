package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"strings"
)

func WindowsToLinuxPath(windowsPath string) string {
	result := strings.ReplaceAll(windowsPath, "\\", separator.Slash)

	if len(result) > 1 && result[1] == ':' {
		driveLetter := strings.ToLower(string(result[0]))
		result = fmt.Sprintf(
			"%s%s%s",
			separator.Slash,
			driveLetter,
			result[2:],
		)
	}

	return result
}
