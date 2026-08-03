package merge_request

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Request) formatTitle(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", r.Title)
	}

	return r.Title
}
