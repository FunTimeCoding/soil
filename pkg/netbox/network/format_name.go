package network

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Interface) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", i.Name)
	}

	return i.Name
}
