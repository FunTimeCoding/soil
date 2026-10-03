package installed

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	lintConstant "github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func changedFiles(
	directory string,
	commit string,
) ([]string, bool) {
	r := run.New()
	r.Panic = false
	r.Directory = directory
	output := r.Start(
		constant.Command,
		constant.Diff,
		constant.NameOnly,
		commit,
		lintConstant.Head,
	)

	if r.Error != nil {
		return nil, false
	}

	return strings.Fields(output), true
}
