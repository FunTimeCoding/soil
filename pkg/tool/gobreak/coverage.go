package gobreak

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
	"slices"
)

func coverage(
	f *break_file.File,
	profile string,
) []string {
	var packages []string

	for _, b := range f.Breaks {
		for _, e := range b.Edits {
			packages = append(
				packages,
				join.Empty(constant.PackagePrefix, filepath.Dir(e.File)),
			)
		}
	}

	if len(packages) == 0 {
		return nil
	}

	slices.Sort(packages)

	return []string{
		constant.CoverFlag,
		profile,
		constant.CoverageFlag,
		join.Comma(slices.Compact(packages)),
	}
}
