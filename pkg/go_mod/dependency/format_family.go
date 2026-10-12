package dependency

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/go_mod/constant"
)

func (d *Dependency) formatFamily(f *option.Format) string {
	if !f.UseColor {
		return d.Family
	}

	switch d.Family {
	case constant.Permissive:
		return consoleConstant.Green("%s", d.Family)
	case constant.Copyleft:
		return consoleConstant.Red("%s", d.Family)
	default:
		return consoleConstant.Yellow("%s", d.Family)
	}
}
