package custom_field

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (f *Field) formatName(o *option.Format) string {
	if o.UseColor {
		return constant.Cyan("%s", f.Name)
	}

	return f.Name
}
