package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"gopkg.in/yaml.v3"
	"maps"
	"path/filepath"
	"slices"
)

func versionedPathConcerns(
	v *virtual_file_system.System,
	path string,
) []*concern.Concern {
	file := filepath.Join(path, "generated", "server", "openapi.yaml")

	if !v.Has(file) {
		return nil
	}

	var s openAPISpec

	if yaml.Unmarshal(v.Read(file), &s) != nil {
		return nil
	}

	var result []*concern.Concern

	for _, route := range slices.Sorted(maps.Keys(s.Paths)) {
		if !constant.VersionedPathPattern.MatchString(route) {
			continue
		}

		result = append(
			result,
			concern.NewPackage(
				constant.VersionedPathKey,
				fmt.Sprintf(
					"versioned API path %s - paths are unversioned (/api/...), break and roll forward",
					route,
				),
				path,
			),
		)
	}

	return result
}
