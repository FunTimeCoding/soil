package service

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
	"strings"
)

func enablementLink(path string) bool {
	directory := filepath.Base(filepath.Dir(path))

	return strings.HasSuffix(directory, constant.WantsSuffix) ||
		strings.HasSuffix(directory, constant.RequiresSuffix)
}
