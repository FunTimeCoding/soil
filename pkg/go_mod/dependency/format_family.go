package dependency

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	go_mod "github.com/funtimecoding/soil/pkg/go_mod/constant"
)

func (d *Dependency) formatFamily(f *option.Format) string {
	if !f.UseColor {
		return d.Family
	}

	switch d.Family {
	case go_mod.Permissive:
		return constant.Green("%s", d.Family)
	case go_mod.Copyleft:
		return constant.Red("%s", d.Family)
	default:
		return constant.Yellow("%s", d.Family)
	}
}
