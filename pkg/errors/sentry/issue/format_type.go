package issue

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/errors/constant"
)

func (i *Issue) formatType(f *option.Format) string {
	if f.UseColor {
		switch i.Type {
		case constant.ErrorType:
			return consoleConstant.Red("%s", i.Type)
		}
	}

	return i.Type
}
