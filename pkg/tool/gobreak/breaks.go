package gobreak

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_file"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/break_result"
	"github.com/funtimecoding/soil/pkg/tool/gobreak/constant"
	"path/filepath"
	"strconv"
)

func breaks(
	root string,
	f *break_file.File,
	directory string,
) ([]*break_result.Result, int, error) {
	profile := filepath.Join(directory, constant.CoverFile)
	b := test(root, f, coverage(f, profile)...)

	if problem := baselineProblem(b); problem != "" {
		return nil, 0, errors.New(problem)
	}

	profiles := coverProfiles(profile)
	modulePath := go_mod.ReadDirectory(root).Module.Mod.Path
	var result []*break_result.Result

	for i, k := range f.Breaks {
		result = append(
			result,
			breakResult(
				root,
				modulePath,
				f,
				k,
				filepath.Join(directory, strconv.Itoa(i)),
				profiles,
			),
		)
	}

	return result, len(b.Passed()), nil
}
