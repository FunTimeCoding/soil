package view

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (v *View) formatTitle(f *option.Format) string {
	title := v.Title

	if title == "" {
		title = NoTitle
	}

	if f.UseColor {
		return console.Cyan("%s", title)
	}

	return title
}
