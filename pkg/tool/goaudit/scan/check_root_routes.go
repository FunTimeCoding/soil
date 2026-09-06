package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"gopkg.in/yaml.v3"
	"maps"
	"path/filepath"
	"slices"
	"strings"
)

func (s *Service) checkRootRoutes(
	v *virtual_file_system.System,
	path string,
) {
	file := filepath.Join(path, "generated", "server", "openapi.yaml")

	if !v.Has(file) {
		return
	}

	var spec openAPISpec

	if yaml.Unmarshal(v.Read(file), &spec) != nil {
		return
	}

	var roots []string

	for _, route := range slices.Sorted(maps.Keys(spec.Paths)) {
		if !strings.HasPrefix(route, webConstant.InterfacePath) {
			roots = append(roots, route)
		}
	}

	if len(roots) == 0 {
		return
	}

	patterns := mountPatterns(v, path, pathConstants(v, path))

	for _, route := range roots {
		if !coveredRoute(patterns, route) {
			s.addConcern(
				constant.RootRouteKey,
				fmt.Sprintf(
					"spec path %s is outside /api/ and no mount pattern covers it",
					route,
				),
				path,
			)
		}
	}
}
