package cluster_type

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Type) formatName(f *option.Format) string {
	if f.UseColor {
		return consoleConstant.Cyan("%s", t.Name)
	}

	return t.Name
}
