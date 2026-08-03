package project

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Project) formatName(f *option.Format) string {
	result := p.CombinedName()

	if f.UseColor {
		result = constant.Cyan("%s", result)
	}

	return result
}
