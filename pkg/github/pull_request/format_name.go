package pull_request

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Request) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", r.Name)
	}

	return r.Name
}
