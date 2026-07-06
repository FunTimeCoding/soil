package issue

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) formatAge(f *option.Format) string {
	result := units.HumanDuration(i.Age())
	color := i.AgeColor()

	if f.UseColor && color != nil {
		result = color(result)
	}

	return result
}
