package merge_request

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Request) formatTitle(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", r.Title)
	}

	return r.Title
}
