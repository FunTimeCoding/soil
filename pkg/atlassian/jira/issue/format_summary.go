package issue

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatSummary(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", i.Summary)
	}

	return i.Summary
}
