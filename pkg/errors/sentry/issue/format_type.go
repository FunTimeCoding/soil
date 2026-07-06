package issue

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/errors/sentry/constant"
)

func (i *Issue) formatType(f *option.Format) string {
	if f.UseColor {
		switch i.Type {
		case constant.ErrorType:
			return console.Red("%s", i.Type)
		}
	}

	return i.Type
}
