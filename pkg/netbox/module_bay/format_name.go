package module_bay

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (b *Bay) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", b.Name)
	}

	return b.Name
}
