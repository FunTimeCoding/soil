package device_type

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Type) formatModel(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", t.Model)
	}

	return t.Model
}
