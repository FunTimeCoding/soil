package custom_field

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (f *Field) formatName(o *option.Format) string {
	if o.UseColor {
		return consoleConstant.Cyan("%s", f.Name)
	}

	return f.Name
}
