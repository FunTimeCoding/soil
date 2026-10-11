package module

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (m *Module) formatConcern(f *option.Format) string {
	if len(m.concern) == 0 {
		return ""
	}

	result := join.Comma(m.concern)

	if f.UseColor {
		result = constant.Yellow("%s", result)
	}

	return result
}
