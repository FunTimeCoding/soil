package branch

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (b *Branch) formatMerged(f *option.Format) string {
	var result string

	if b.Merged {
		result = Merged

		if f.UseColor {
			return console.Green("%s", result)
		}
	} else {
		result = Unmerged

		if f.UseColor {
			return console.Yellow("%s", result)
		}
	}

	return result
}
