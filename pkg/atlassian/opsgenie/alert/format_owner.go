package alert

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (a *Alert) formatOwner(f *option.Format) string {
	var result string

	if a.Owner != "" {
		result = a.shortenUser(a.Owner)

		if f.UseColor {
			result = consoleConstant.Green("%s", result)
		}
	} else if !f.HasTag(consoleConstant.TagDense) {
		result = constant.OpsgenieNoOwner

		if f.UseColor {
			result = consoleConstant.Red("%s", result)
		}
	}

	return result
}
