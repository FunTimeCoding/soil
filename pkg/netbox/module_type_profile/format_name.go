package module_type_profile

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Profile) formatName(f *option.Format) string {
	if f.UseColor {
		return constant.Cyan("%s", p.Name)
	}

	return p.Name
}
