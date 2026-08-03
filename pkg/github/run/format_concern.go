package run

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (r *Run) formatConcern(f *option.Format) string {
	if len(r.concern) == 0 {
		return ""
	}

	result := join.Comma(r.concern)

	if f.UseColor {
		result = constant.Yellow("%s", result)
	}

	return result
}
