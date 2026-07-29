package issue

import (
	"github.com/docker/go-units"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) FormatAge(f *option.Format) string {
	if i.Create.IsZero() {
		return constant.JiraNoAge
	}

	result := units.HumanDuration(i.Age())

	if f.UseColor && i.ageColor != nil {
		return i.ageColor(result)
	}

	return result
}
