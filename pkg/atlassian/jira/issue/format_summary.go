package issue

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatSummary(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", i.Summary)
	}

	return i.Summary
}
