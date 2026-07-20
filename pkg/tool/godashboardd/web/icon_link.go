package web

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"strings"
)

func iconLink(icon string) string {
	if strings.HasPrefix(icon, "http") {
		return icon
	}

	return join.Empty(constant.IconHost, icon)
}
