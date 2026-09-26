package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func aliasPath(path string) string {
	if strings.HasPrefix(path, constant.MergedPrefix) {
		return strings.TrimPrefix(path, constant.UsrDirectory)
	}

	if strings.HasPrefix(path, constant.UnmergedPrefix) {
		return join.Empty(constant.UsrDirectory, path)
	}

	return ""
}
