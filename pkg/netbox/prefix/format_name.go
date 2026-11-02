package prefix

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Prefix) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", p.Name)
	}

	return p.Name
}
