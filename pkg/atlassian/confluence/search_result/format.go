package search_result

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Result) Format(f *option.Format) string {
	return status.New(f).String(r.Name).Format()
}
