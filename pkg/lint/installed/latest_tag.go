package installed

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func latestTag(directory string) string {
	r := run.New()
	r.Panic = false
	r.Directory = directory
	output := r.Start(
		constant.Command,
		constant.Describe,
		constant.Tags,
		constant.NoAbbreviation,
	)

	if r.Error != nil {
		return ""
	}

	return strings.TrimSpace(output)
}
