package changed

import (
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func Unstaged(directory string) []string {
	c := run.New()
	c.Directory = directory
	c.Start(constant.Command, constant.Diff, constant.NameOnly)

	return slice.StripEmpty(split.NewLine(strings.TrimSpace(c.OutputString)))
}
