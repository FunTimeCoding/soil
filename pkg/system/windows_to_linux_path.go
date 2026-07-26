package system

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func WindowsToLinuxPath(windowsPath string) string {
	result := strings.ReplaceAll(windowsPath, "\\", constant.Slash)

	if len(result) > 1 && result[1] == ':' {
		driveLetter := strings.ToLower(string(result[0]))
		result = fmt.Sprintf(
			"%s%s%s",
			constant.Slash,
			driveLetter,
			result[2:],
		)
	}

	return result
}
