package changed

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/git/constant"
	"github.com/funtimecoding/soil/pkg/strings/slice"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"github.com/funtimecoding/soil/pkg/system/run"
	"strings"
)

func (r *Range) Files(directory string) []string {
	if r.All || r.None {
		return nil
	}

	c := run.New()
	c.Directory = directory

	if r.Staged {
		c.Start(
			constant.Command,
			constant.Diff,
			constant.Cached,
			constant.NameOnly,
		)
	} else {
		c.Start(
			constant.Command,
			constant.Diff,
			constant.NameOnly,
			fmt.Sprintf("%s..%s", r.Base, r.Head),
		)
	}

	return slice.StripEmpty(split.NewLine(strings.TrimSpace(c.OutputString)))
}
