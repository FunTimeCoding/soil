package alert

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatOwner(f *option.Format) string {
	var result string

	if a.Owner != "" {
		result = a.shortenUser(a.Owner)

		if f.UseColor {
			result = console.Green("%s", result)
		}
	} else if !f.HasTag(constant.TagDense) {
		result = NoOwner

		if f.UseColor {
			result = console.Red("%s", result)
		}
	}

	return result
}
