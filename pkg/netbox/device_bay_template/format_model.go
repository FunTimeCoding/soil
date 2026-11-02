package device_bay_template

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Template) formatModel(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", t.Model)
	}

	return t.Model
}
