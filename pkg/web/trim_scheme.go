package web

import (
	"github.com/funtimecoding/soil/pkg/web/constant"
	"strings"
)

func TrimScheme(s string) string {
	if strings.HasPrefix(s, constant.SecurePrefix) {
		return strings.TrimPrefix(s, constant.SecurePrefix)
	}

	if strings.HasPrefix(s, constant.InsecurePrefix) {
		return strings.TrimPrefix(s, constant.InsecurePrefix)
	}

	return s
}
