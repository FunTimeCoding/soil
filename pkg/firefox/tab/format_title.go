package tab

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/firefox/constant"
)

func (t *Tab) formatTitle(f *option.Format) string {
	title := t.Title

	if title == "" {
		title = constant.NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
