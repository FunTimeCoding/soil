package issue

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatSummary(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", i.Summary)
	}

	return i.Summary
}
