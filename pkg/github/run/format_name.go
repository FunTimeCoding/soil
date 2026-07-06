package run

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Run) formatName(f *option.Format) string {
	result := r.Repository.FullName

	if f.UseColor {
		result = console.Cyan("%s", result)
	}

	return result
}
