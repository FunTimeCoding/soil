package virtual_machine

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Machine) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", r.Name)
	}

	return r.Name
}
