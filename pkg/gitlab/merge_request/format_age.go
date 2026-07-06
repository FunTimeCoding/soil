package merge_request

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (r *Request) formatAge(f *option.Format) string {
	result := units.HumanDuration(r.Age())

	if f.UseColor && r.ageColor != nil {
		result = r.ageColor(result)
	}

	return result
}
