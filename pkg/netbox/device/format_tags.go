package device

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"slices"
)

func (d *Device) formatTags(f *option.Format) string {
	if len(d.Tags) == 0 {
		return ""
	}

	var toFilter []string

	for _, a := range f.Filters {
		if a.Key == constant.TagKey {
			toFilter = append(toFilter, a.Value)
		}
	}

	var filtered []string

	for _, a := range d.Tags {
		if !slices.Contains(toFilter, a) {
			filtered = append(filtered, a)
		}
	}

	if len(filtered) == 0 {
		return ""
	}

	result := join.Comma(filtered)

	if f.UseColor {
		result = console.Yellow("%s", result)
	}

	return result
}
