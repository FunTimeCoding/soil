package system

import (
	"fmt"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func WindowsToLinuxPath(windowsPath string) string {
	result := strings.ReplaceAll(windowsPath, "\\", stringsConstant.Slash)

	if len(result) > 1 && result[1] == ':' {
		driveLetter := strings.ToLower(string(result[0]))
		result = fmt.Sprintf(
			"%s%s%s",
			stringsConstant.Slash,
			driveLetter,
			result[2:],
		)
	}

	return result
}
