package service

import (
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func platformSuffix(stem string) string {
	for _, suffix := range lintConstant.PlatformSuffixes {
		if strings.HasSuffix(stem, suffix) {
			return suffix
		}
	}

	return ""
}
