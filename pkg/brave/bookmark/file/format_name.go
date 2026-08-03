package file

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (n *Node) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", n.Name)
	}

	return n.Name
}
