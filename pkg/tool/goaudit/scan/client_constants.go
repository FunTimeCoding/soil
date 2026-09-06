package scan

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"maps"
	"path/filepath"
	"slices"
	"strings"
)

func ClientConstants(v *virtual_file_system.System) []*concern.Concern {
	directories := make(map[string]bool)

	for _, path := range v.Files() {
		segments := strings.Split(path, "/")

		if isTestdataPath(segments) || len(segments) < 2 {
			continue
		}

		if segments[len(segments)-2] == constant.ConstantDirectory {
			directories[filepath.Dir(path)] = true
		}
	}

	var result []*concern.Concern

	for _, directory := range slices.Sorted(maps.Keys(directories)) {
		result = append(result, clientConstantConcerns(v, directory)...)
	}

	return result
}
