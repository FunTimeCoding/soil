package branch

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
)

func (b *Branch) formatMerged(f *option.Format) string {
	var result string

	if b.Merged {
		result = constant.BranchMerged

		if f.UseColor {
			return console.Green("%s", result)
		}
	} else {
		result = constant.BranchUnmerged

		if f.UseColor {
			return console.Yellow("%s", result)
		}
	}

	return result
}
