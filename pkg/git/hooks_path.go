package git

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func HooksPath(directory string) string {
	r := run.New()
	r.Panic = false
	r.Directory = directory
	r.Start(
		constant.Command,
		constant.Configuration,
		constant.Get,
		constant.HooksPathKey,
	)

	return strings.TrimSpace(r.OutputString)
}
