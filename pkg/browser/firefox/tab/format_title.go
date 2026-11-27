package tab

import (
	"github.com/funtimecoding/soil/pkg/browser/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Tab) formatTitle(f *option.Format) string {
	title := t.Title

	if title == "" {
		title = constant.FirefoxNoTitle
	}

	if f.UseColor {
		return consoleConstant.Cyan("%s", title)
	}

	return title
}
