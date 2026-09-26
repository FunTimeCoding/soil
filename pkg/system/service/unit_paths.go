package service

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"path/filepath"
)

func unitPaths() map[string]string {
	result := make(map[string]string)

	for _, directory := range constant.SystemdDirectories {
		if !system.DirectoryExists(directory) {
			continue
		}

		for _, name := range system.FilesByExtension(
			directory,
			constant.UnitSuffix,
		) {
			if enablementLink(name) {
				continue
			}

			unit := filepath.Base(name)

			if _, okay := result[unit]; okay {
				continue
			}

			result[unit] = resolve(name)
		}
	}

	return result
}
