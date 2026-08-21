package file_identity

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func trimPlatformSuffix(stem string) string {
	for _, suffix := range constant.PlatformSuffixes {
		if strings.HasSuffix(stem, suffix) {
			return strings.TrimSuffix(stem, suffix)
		}
	}

	return stem
}
