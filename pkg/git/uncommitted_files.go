package git

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
)

func UncommittedFiles(directory string) ([]string, error) {
	r := run.New().NoPanic()
	r.Directory = directory
	r.Start(
		constant.Command,
		constant.Diff,
		constant.NameOnly,
		constant.Relative,
		constant.HeadReference,
		constant.Pathspec,
		library.CurrentDirectory,
	)

	if r.Error != nil {
		return nil, r.Error
	}

	var result []string

	for _, line := range split.NewLine(r.OutputString) {
		if line != "" {
			result = append(result, line)
		}
	}

	return result, nil
}
