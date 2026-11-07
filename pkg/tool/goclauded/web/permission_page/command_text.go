package permission_page

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
)

func commandText(command string) string {
	runes := []rune(command)

	if len(runes) <= constant.PermissionCommandLength {
		return command
	}

	return join.Empty(string(runes[:constant.PermissionCommandLength]), "…")
}
