package tunnel

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Tunnel) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", t.Name)
	}

	return t.Name
}
