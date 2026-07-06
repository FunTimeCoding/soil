package branch

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (b *Branch) formatAge(f *option.Format) string {
	result := units.HumanDuration(b.Age())

	if f.UseColor && b.ageColor != nil {
		result = b.ageColor(result)
	}

	return result
}
