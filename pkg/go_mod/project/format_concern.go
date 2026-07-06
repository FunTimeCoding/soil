package project

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (p *Project) formatConcern(f *option.Format) string {
	if len(p.concern) == 0 {
		return ""
	}

	result := join.Comma(p.concern)

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
