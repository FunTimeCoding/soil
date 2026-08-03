package inventory_item

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Item) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", i.Name)
	}

	return i.Name
}
