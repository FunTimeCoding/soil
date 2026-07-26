package view

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/sublime/constant"
)

func (v *View) formatTitle(f *option.Format) string {
	title := v.Title

	if title == "" {
		title = constant.NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
