package scan

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"sort"
	"strings"
)

func ConstantPlacement(v *virtual_file_system.System) []*concern.Concern {
	var result []*concern.Concern
	directories := make(map[string]bool)
	nested := make(map[string]bool)

	for _, path := range v.Files() {
		segments := strings.Split(path, "/")

		if isTestdataPath(segments) {
			continue
		}

		last := len(segments) - 1

		if segments[last] == constant.ConstantFileName &&
			(last == 0 || segments[last-1] != constant.ConstantDirectory) {
			result = append(
				result,
				concern.NewFile(
					constant.ConstantFileKey,
					constant.ConstantFileText,
					path,
					false,
				),
			)
		}

		for i, s := range segments[:last] {
			if s != constant.ConstantDirectory {
				continue
			}

			directories[strings.Join(segments[:i+1], "/")] = true

			if last-i > 1 && !isTestHome(segments[i+1]) {
				nested[strings.Join(segments[:i+2], "/")] = true
			}
		}
	}

	for directory := range directories {
		if isConstantRoot(directory) {
			continue
		}

		result = append(
			result,
			concern.NewFile(
				constant.ConstantDepthKey,
				constant.ConstantDepthText,
				directory,
				false,
			),
		)
	}

	for directory := range nested {
		result = append(
			result,
			concern.NewFile(
				constant.ConstantNestedKey,
				constant.ConstantNestedText,
				directory,
				false,
			),
		)
	}

	sort.Slice(
		result,
		func(i, j int) bool {
			if result[i].Path == result[j].Path {
				return result[i].Key < result[j].Key
			}

			return result[i].Path < result[j].Path
		},
	)

	return result
}
