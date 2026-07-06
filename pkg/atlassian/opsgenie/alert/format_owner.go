package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (a *Alert) formatOwner(f *option.Format) string {
	var result string

	if a.Owner != "" {
		result = a.shortenUser(a.Owner)

		if f.UseColor {
			result = console.Green("%s", result)
		}
	} else if !f.HasTag(tag.Dense) {
		result = NoOwner

		if f.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
