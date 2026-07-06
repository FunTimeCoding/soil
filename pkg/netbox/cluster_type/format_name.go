package cluster_type

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Type) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", t.Name)
	}

	return t.Name
}
