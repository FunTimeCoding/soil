package virtual_network

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (n *Network) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", n.Name)
	}

	return n.Name
}
