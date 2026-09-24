package source

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func IsGeneratedHeader(content string) bool {
	for _, line := range strings.SplitN(
		content,
		stringsConstant.Unix,
		constant.GeneratedProbe,
	) {
		trimmed := strings.TrimSpace(line)

		if !strings.HasPrefix(trimmed, stringsConstant.DoubleSlash) && trimmed != "" {
			return false
		}

		if strings.HasPrefix(trimmed, constant.GeneratedMarker) &&
			strings.HasSuffix(trimmed, constant.GeneratedSuffix) {
			return true
		}
	}

	return false
}
