package view

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (v *View) formatDirty(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Yellow("modified")
	}

	return "modified"
}
