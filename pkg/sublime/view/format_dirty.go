package view

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (v *View) formatDirty(f *option.Format) string {
	if f.UseColor {
		return constant.Yellow("modified")
	}

	return "modified"
}
