package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/relocation"
	"path/filepath"
)

func soleMoveTarget(
	plan *relocation.Plan,
	filename string,
) string {
	result := ""

	for _, entry := range plan.Entries {
		if plan.Set.Position(entry.File.Pos()).Filename != filename {
			continue
		}

		target := filepath.Join(plan.MoveDirectory, entry.TargetFile)

		if result != "" && result != target {
			return ""
		}

		result = target
	}

	return result
}
