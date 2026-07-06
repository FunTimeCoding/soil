package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"strings"
)

func iconLink(icon string) string {
	if strings.HasPrefix(icon, "http") {
		return icon
	}

	return join.Empty(iconHost, icon)
}
