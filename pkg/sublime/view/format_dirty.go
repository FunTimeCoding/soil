package view

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (v *View) formatDirty(f *option.Format) string {
	if f.UseColor {
		return console.Yellow("modified")
	}

	return "modified"
}
