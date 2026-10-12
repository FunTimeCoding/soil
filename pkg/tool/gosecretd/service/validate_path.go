package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/constant"
	"slices"
	"strings"
)

func validatePath(path string) error {
	if path == "" {
		return validation.New("path is required")
	}

	if strings.HasPrefix(path, constant.Folder) ||
		strings.HasSuffix(path, constant.Folder) {
		return validation.New("path starts or ends with a slash: %s", path)
	}

	if slices.Contains(
		strings.Split(path, constant.Folder),
		constant.ParentSegment,
	) {
		return validation.New("path climbs out of the store: %s", path)
	}

	return nil
}
