package service

import (
	"fmt"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func targetDirectory(
	source *packages.Package,
	target *packages.Package,
	targetPackagePath string,
) (string, error) {
	if target != nil {
		if len(target.GoFiles) == 0 {
			return "", fmt.Errorf(
				"target package has no Go files: %s",
				targetPackagePath,
			)
		}

		return filepath.Dir(target.GoFiles[0]), nil
	}

	if source.Module == nil {
		return "", fmt.Errorf("source package has no module information")
	}

	prefix := fmt.Sprintf("%s/", source.Module.Path)

	if !strings.HasPrefix(targetPackagePath, prefix) {
		return "", fmt.Errorf(
			"target package is outside the module: %s",
			targetPackagePath,
		)
	}

	relative := strings.TrimPrefix(targetPackagePath, prefix)

	return filepath.Join(source.Module.Dir, relative), nil
}
