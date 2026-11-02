package merge_request

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Request) formatTitle(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", r.Title)
	}

	return r.Title
}
