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

func runBreaks(
	root string,
	f *break_file.File,
	directory string,
) ([]*break_result.Result, int, error) {
	if problem := expectProblem(f); problem != "" {
		return nil, 0, errors.New(problem)
	}

	profile := filepath.Join(directory, constant.CoverFile)
	b := runTests(root, f, coverage(f, profile)...)

	if problem := baselineProblem(b); problem != "" {
		return nil, 0, errors.New(problem)
	}

	profiles := coverProfiles(profile)
	modulePath := go_mod.ReadDirectory(root).Module.Mod.Path
	var result []*break_result.Result

	for i, k := range f.Breaks {
		result = append(
			result,
			runBreak(
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
